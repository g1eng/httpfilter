package auth

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//Provider represents authentication Provider for httpfilter.
type Provider interface {
	Authenticate(handle http.HandlerFunc, credentials ...string) http.HandlerFunc
	RouterAuthenticate(handle httprouter.Handle, credentials ...string) httprouter.Handle
	Log(string)
}
