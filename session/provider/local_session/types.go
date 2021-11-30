package local_session

import (
	"crypto/x509"
	"io"
	"math/rand"
)

//BasicStore is session store for basic authentication.
type BasicStore struct {
	userSessions    map[string]string //userSessions holds authenticated sessions for ensuring that an account can sign in with only one device
	userCredentials map[string]string //userCredentials holds plain-text credentials for basic authentication for all user within a slice of string
	Logger          []io.Writer       //logging interface
}

func NewBasicStore(credential map[string]string, logWriters ...io.Writer) BasicStore {
	return BasicStore{
		userSessions:    map[string]string{},
		userCredentials: credential,
		Logger:          logWriters,
	}
}

//TokenStore is session token storage with in-memory data.
//That holds uniq server session id for the LDAP connection, and
//a map of client session which is keyed by hashed set of ip address
//and client user agent string.
//
type TokenStore struct {
	serverSessionId uint64            //serverSessionId is unique id calculated at the start
	userSessions    map[string]string //userSessions holds session secret keyed by client IP + UA
	customHeader    string            //Custom header name for HeaderAuth
	accessToken     string            //protection secret for additional user validation. GetAccessToken returns this value.
	Logger          []io.Writer       //logging interface
}

//NewLocalSession generates new TokenStore instance and return it.
func NewLocalSession(customHeader string, logWriters ...io.Writer) TokenStore {
	return TokenStore{
		serverSessionId: rand.Uint64(),
		userSessions:    map[string]string{},
		customHeader:    customHeader,
		accessToken:     CalcBs32(),
		Logger:          logWriters,
	}
}

//SimpleLoginCredential is JSON wrapper which stores login request from
//clients.
type SimpleLoginCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SimpleMFACredential struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	OTPSecret string `json:"secret"`
}

type SingleOTPCredential struct {
	OTPSecret string `json:"secret"`
}

type CertificateCredential struct {
	Cert *x509.Certificate `json:"certificate"`
}

type OAuthCredential struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
