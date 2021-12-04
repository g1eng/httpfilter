package basic

import (
	"io"
	"os"
	"strings"
)

type Authenticator struct {
	userCredentials map[string]string
	Logger          []io.Writer
}

func NewBasicAuth(cred string, w ...io.Writer) *Authenticator {
	auth := &Authenticator{
		Logger: w,
	}
	var c map[string]string
	f, err := os.Open(cred)
	if err != nil {
		c, err = ParseHTPasswd(strings.NewReader(cred))
		if err != nil {
			auth.Log(err.Error())
			os.Exit(1)
		}
		auth.userCredentials = c
	} else {
		c, err = ParseHTPasswd(f)
		if err != nil {
			auth.Log(err.Error())
			os.Exit(1)
		}
	}
	return auth
}
