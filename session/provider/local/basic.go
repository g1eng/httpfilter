package local

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/g1eng/httpfilter/session/responder"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strings"
)

// SplitBasicCred splits input string with colon and validate the result length is 2.
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
// Picked and modified from https://github.com/distribution/distribution/blob/v2.7.1/registry/auth/htpasswd/htpasswd.go
func ParseHTPasswd(rd io.Reader) (map[string]string, error) {
	entries := map[string]string{}
	scanner := bufio.NewScanner(rd)
	var line int
	for scanner.Scan() {
		line++ // 1-based line numbering
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

		entries[t[:i]] = t[i+1:]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

// Log provides internal logging mechanism for TokenStore
func (sess BasicStore) Log(logLine string) (errorArray []error) {
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

//For compatibility and no real action
func (sess *BasicStore) RegisterClient(_ *http.Request) (string, error) {
	return "", nil
}

//For compatibility and no real action
func (sess *BasicStore) RevokeClient(_ *http.Request) error {
	return nil
}

//For compatibility and no real action
func (sess *BasicStore) GetAccessToken() string {
	return ""
}

//Auth handles authentication process via Basic authentication.
//If any unauthorized access, it sends WWW-Authenticate header for the client.
func (sess *BasicStore) Auth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(authHeader) != 2 || strings.ToLower(authHeader[0]) != "basic" {
			w.Header().Set("WWW-Authenticate", `Basic realm="basic authentication"`)
			responder.Write400(w)
			return
		}
		payload, err := base64.StdEncoding.DecodeString(authHeader[1])
		if err != nil {
			responder.Write401(w)
			return
		}

		user, cryptPassword, err := SplitBasicCred(string(payload))

		//Invalid authHeader
		if err != nil {
			responder.Write400(w)
			return
		}

		for u, c := range sess.userCredentials {
			if u == user && c == cryptPassword {
				handle(w, r, ps)
				return
			}
		}
		responder.Write401(w)
		return
	}
}
