package provider

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//Provider represents authentication provider.
type Provider interface {
	Auth(handle httprouter.Handle) httprouter.Handle
	RegisterClient(r *http.Request) (string, error)
	RevokeClient(r *http.Request) error
	GetAccessToken() string     //GetAccessToken returns temporary access token for protected resources
	Log(logLine string) []error //write log with preset log writer
}
