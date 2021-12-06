package session

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//Provider represents authentication provider.
type Provider interface {
	Authorize(handle httprouter.Handle, _ ...string) httprouter.Handle
	RegisterClient(r *http.Request) (string, error)
	RevokeClient(r *http.Request) error
	GetAccessToken() string     //GetAccessToken returns temporary access token for protected resources
	Log(logLine string) []error //write log with preset log writer
}
