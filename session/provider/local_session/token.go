package local_session

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/g1eng/w3fs/middleware/filter/session/responder"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Log provides internal logging mechanism for TokenStore
func (sess TokenStore) Log(logLine string) (errorArray []error) {
	if len(sess.Logger) == 0 {
		log.Println(logLine)
		return errorArray
	}
	for _, logger := range sess.Logger {
		if _, err := logger.Write([]byte(logLine)); err != nil {
			errorArray = append(errorArray, err)
		}
	}
	return errorArray
}

// RegisterClient generates new secret and registers client and its secret
// in TokenStore.
func (sess TokenStore) RegisterClient(r *http.Request) (string, error) {
	key := calcSessionKey(r)
	if key == "" {
		return "", errors.New("invalid client: pair of IP and User Agent not found")
	} else {
		bs := CalcBs160()
		sess.userSessions[key] = bs
		sess.Log("new secret registered for " + key)
		//log.Println("new secret", sess.userSessions[key])
		return sess.userSessions[key], nil
	}
}

// RevokeClient revokes registered client for the session.
// This method removes registered credential for the client
func (sess TokenStore) RevokeClient(r *http.Request) error {
	key := calcSessionKey(r)
	if sess.userSessions[key] == "" {
		return errors.New("client not registered")
	} else {
		delete(sess.userSessions, key)
		return nil
	}
}

// Auth provides local authentication handler for httprouter.Handle handlers.
// This method searches client information in TokenStore and matches header
// value with the secret registered in userSessions.
// If the matched result is success, it provides protected route for the client.
//
//[Note] customHeader is used for authentication.
func (sess TokenStore) Auth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		key := calcSessionKey(r)
		val := sess.userSessions[key]
		if val == "" {
			responder.Write403(w)
		} else if r.Header.Get(sess.customHeader) == "" {
			sess.Log(fmt.Sprintf("%v %v %v %v %d // blank token", r.Method, r.RemoteAddr, r.Referer(), r.UserAgent(), http.StatusUnauthorized))
			responder.Write403(w)
		} else if r.Header.Get(sess.customHeader) != val {
			sess.Log(fmt.Sprintf("%v %v %v %v %d // invalid token", r.Method, r.RemoteAddr, r.Referer(), r.UserAgent(), http.StatusUnauthorized))
			responder.Write403(w)
		} else {
			sess.Log(fmt.Sprintf("%v %v %v %v %d // -", r.Method, r.RemoteAddr, r.Referer(), r.UserAgent(), http.StatusOK))
			h(w, r, ps)
		}
	}
}

func (sess *TokenStore) GetAccessToken() string {
	return sess.accessToken
}

func (sess *TokenStore) RegenerateToken() string {
	sess.accessToken = CalcBs32()
	return sess.accessToken
}

func (sess *TokenStore) ValidateToken(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var c SimpleMFACredential
		if r.ContentLength > 1024 {
			responder.Write403(w)
			return
		}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&c); err != nil {
			sess.Log(fmt.Sprintf("%v %v %v %v %d // json decode failure", r.Method, r.RemoteAddr, r.Referer(), r.UserAgent(), http.StatusForbidden))
			responder.Write403(w)
		} else if c.OTPSecret == "" {
			sess.Log(fmt.Sprintf("%v %v %v %v %d // blank access token", r.Method, r.RemoteAddr, r.Referer(), r.UserAgent(), http.StatusForbidden))
			responder.Write403(w)
		} else if c.OTPSecret != sess.accessToken {
			sess.Log(fmt.Sprintf("%v %v %v %v %d // invalid access token", r.Method, r.RemoteAddr, r.Referer(), r.UserAgent(), http.StatusForbidden))
			responder.Write403(w)
		} else {
			handle(w, r, ps)
		}
	}
}
