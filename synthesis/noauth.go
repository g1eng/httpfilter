package synthesis

import (
	"github.com/g1eng/httpfilter/session/responder"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// NoAuth is a dummy function to bypass httprouter.Handle.
// It is useful to route unprotected resources.
func NoAuth(h httprouter.Handle) httprouter.Handle {
	return h
}

// False is special function to disable specific httprouter.Handle
// on a route. It is useful to programmatically shutoff access to
// the specific resource.
func False(_ httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		responder.Write403(w)
	}
}

// PlainNoAuth is a dummy function to bypass http.HandlerFunc.
// It is useful to route unprotected resources.
func PlainNoAuth(h http.HandlerFunc) http.HandlerFunc {
	return h
}

// PlainFalse is special function to disable specific http.HandlerFunc
// on a route. It is useful to programmatically shutoff access to
// the specific resource.
func PlainFalse(_ http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		responder.Write403(w)
	}
}
