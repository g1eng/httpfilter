package token

import (
	"io"
	"math/rand"
)

//Store is session token storage with in-memory data.
//That holds uniq server session id for the LDAP connection, and
//a map of client session which is keyed by hashed set of ip address
//and client user agent string.
//
type Store struct {
	serverSessionId uint64            //serverSessionId is unique id calculated at the start
	userSessions    map[string]string //userSessions holds session secret keyed by client IP + UA
	customHeader    string            //Custom header name for HeaderAuth
	accessToken     string            //protection secret for additional user validation. GetAccessToken returns this value.
	Logger          []io.Writer       //logging interface
}

//NewStore generates new token.Store instance, which represents in-memory temporary token store.
//FIXME: implement salt and regenerate accessToken in effective sequence (against rainbow table attack)
func NewStore(customHeader string, logWriters ...io.Writer) Store {
	return Store{
		serverSessionId: rand.Uint64(),
		userSessions:    map[string]string{},
		customHeader:    customHeader,
		accessToken:     CalcBs32(),
		Logger:          logWriters,
	}
}
