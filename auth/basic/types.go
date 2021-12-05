package basic

import (
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
	"strings"
)

//Authenticator is auth.Provider implementation for basic authentication.
type Authenticator struct {
	userCredentials map[string][]byte //userCredentials specified in a map of string to byte arrays
	Logger          []io.Writer       //log writer set
}

//NewBasicAuth generates new filter instance for basic authentication.
func NewBasicAuth(cred string, logger ...io.Writer) (*Authenticator, error) {
	auth := &Authenticator{
		Logger: logger,
	}
	var c map[string][]byte

	//try to open `cred` as a file
	f, err := os.Open(cred)
	if err != nil {
		c, err = ParseHTPasswd(strings.NewReader(cred))
		if err != nil {
			auth.Log(err.Error())
			return nil, err
		}
		for k, v := range c {
			c[k], err = bcrypt.GenerateFromPassword(v, 12)
			if err != nil {
				return nil, err
			}
		}
		auth.userCredentials = c
	} else {
		log.Println("isfile")
		auth.userCredentials, err = ParseHTPasswd(f)
		if err != nil {
			auth.Log(err.Error())
			return nil, err
		}
	}
	return auth, nil
}
