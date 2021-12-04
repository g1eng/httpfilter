package basic

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
	"os"
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
// Cherry-picked and modified from https://github.com/distribution/distribution/blob/v2.7.1/registry/auth/htpasswd/htpasswd.go
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

func (b *Authenticator) Authenticate(_ http.HandlerFunc, _ ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//stub
	}
}

//RouterAuthenticate handles authentication process via Basic authentication.
//If any unauthorized access, it sends WWW-RouterAuthenticate header for the client
//and terminate the session.
func (b *Authenticator) RouterAuthenticate(handle httprouter.Handle, _ ...string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(authHeader) != 2 || strings.ToLower(authHeader[0]) != "basic" {
			w.Header().Set("WWW-RouterAuthenticate", `Basic realm="basic authentication"`)
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
		for u, c := range b.userCredentials {
			if u == user && c == cryptPassword {
				handle(w, r, ps)
				return
			}
		}
		responder.Write401(w)
		return
	}
}
