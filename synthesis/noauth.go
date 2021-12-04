package synthesis

import (
	"github.com/g1eng/httpfilter/session/responder"
	"net/http"
)

// UnAuth is special function to disable specific route with 401
// status code. It is useful to programmatically shutoff access to
// the specific resource.
func UnAuth(_ http.HandlerFunc, _ ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		responder.Write401(w)
	}
}

// NoAuth is a dummy function to bypass http.HandlerFunc.
// It is useful to route unprotected resources.
func NoAuth(h http.HandlerFunc, _ ...string) http.HandlerFunc {
	return h
}

// Forbid is special function to disable specific http.HandlerFunc
// on a route. It is useful to programmatically shutoff access to
// the specific resource.
func Forbid(_ http.HandlerFunc, _ ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		responder.Write403(w)
	}
}
