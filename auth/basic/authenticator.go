package basic

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/g1eng/httpfilter/session/responder"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// SplitBasicCred splits input string with htpasswd format. Input string is
// separated with colon and validated as slice which has length 2.
// If the length of result of split is not 2, it returns error.
func SplitBasicCred(c string) (user string, password string, err error) {
	credential := strings.SplitN(c, ":", 2)

	//Invalid authHeader
	if len(credential) != 2 {
		return "", "", errors.New("invalid credential format")
	}
	return credential[0], credential[1], nil
}

// ParseHTPasswd parses the contents of htpasswd. This will read all the
// entries in the file, whether or not they are needed. An error is returned
// if a syntax errors are encountered or if the reader fails.
// Cherry-picked and modified from https://github.com/distribution/distribution/blob/v2.7.1/registry/auth/htpasswd/htpasswd.go
func ParseHTPasswd(rd io.Reader) (map[string][]byte, error) {
	entries := map[string][]byte{}
	scanner := bufio.NewScanner(rd)
	var line int
	for scanner.Scan() {
		line++ // 1-based line numbering
		log.Println("line: ", line)
		t := strings.TrimSpace(scanner.Text())

		if len(t) < 1 {
			continue
		}

		// lines that *begin* with a '#' are considered comments
		if t[0] == '#' {
			continue
		}

		i := strings.Index(t, ":")
		if i < 0 || i >= len(t) {
			return nil, fmt.Errorf("htpasswd: invalid entry at line %d: %q", line, scanner.Text())
		}

		entries[t[:i]] = []byte(t[i+1:])
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	log.Println("entries: ", entries)
	return entries, nil
}

// Log provides internal logging mechanism for Authenticator
func (b Authenticator) Log(logLine string) {
	if len(b.Logger) == 0 {
		log.Println(logLine)
		return
	}
	errorStr := ""
	for _, logger := range b.Logger {
		if _, err := logger.Write([]byte(logLine)); err != nil {
			errorStr = fmt.Sprintf("%v\n", err)
		}
	}
	if errorStr != "" {
		_, _ = fmt.Fprintf(os.Stderr, errorStr)
	}
}

func (b *Authenticator) getAuthPayload(w http.ResponseWriter, r *http.Request) (string, error) {
	authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(authHeader) != 2 || strings.ToLower(authHeader[0]) != "basic" {
		log.Println("invalid payload")
		w.Header().Set("WWW-Authenticate", `Basic realm="basic authentication"`)
		responder.Write400(w)
		return "", fmt.Errorf("invalid header for basic auth %v %v", r.RemoteAddr, r.UserAgent())
	}
	payload, err := base64.StdEncoding.DecodeString(authHeader[1])
	if err != nil {
		log.Println(err)
		responder.Write400(w)
		return "", fmt.Errorf("credential not found for basic auth %v %v %v", r.RemoteAddr, r.UserAgent(), r.Header.Get("Authorization"))
	}
	return string(payload), nil
}

func (b *Authenticator) Authenticate(handler http.HandlerFunc, _ ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := b.getAuthPayload(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		user, plainPass, err := SplitBasicCred(p)

		//Invalid authHeader
		if err != nil {
			log.Println("invalid authorization header")
			responder.Write400(w)
			return
		}
		log.Println("plain password: ", plainPass)
		for u, c := range b.userCredentials {
			if u == user {
				err = bcrypt.CompareHashAndPassword(c, []byte(plainPass))
				//cost, _ := bcrypt.Cost(c)
				//log.Println("cost for ", u, cost)
				log.Printf("user %s, err: %v", u, err)
				if err == nil {
					log.Println("authenticated")
					handler(w, r)
					return
				}
			}
		}
		responder.Write401(w)
		return
	}
}

//RouterAuthenticate handles authentication process via Basic authentication.
//If any unauthorized access, it sends WWW-RouterAuthenticate header for the client
//and terminate the session.
func (b *Authenticator) RouterAuthenticate(handle httprouter.Handle, _ ...string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		p, err := b.getAuthPayload(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		user, plainPass, err := SplitBasicCred(p)

		//Invalid authHeader
		if err != nil {
			responder.Write400(w)
			return
		}
		for u, c := range b.userCredentials {
			if u == user {
				err = bcrypt.CompareHashAndPassword(c, []byte(plainPass))
				log.Printf("user %s, err: %v", u, err)
				if err == nil {
					handle(w, r, ps)
					return
				}
			}
		}

		responder.Write401(w)
		return
	}
}
