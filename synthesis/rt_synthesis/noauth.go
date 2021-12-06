package rt_synthesis

import (
	"github.com/g1eng/httpfilter/session/responder"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// NoAuth is a dummy function to bypass httprouter.Handle.
// It is useful to route unprotected resources.
func NoAuth(h httprouter.Handle, _ ...string) httprouter.Handle {
	return h
}

// Forbid is special function to disable specific route with 403
// status code. It is useful to programmatically shutoff access to
// the specific resource.
func Forbid(_ httprouter.Handle, _ ...string) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		responder.Write403(w)
	}
}

// UnAuth is special function to disable specific route with 401
// status code. It is useful to programmatically shutoff access to
// the specific resource.
func UnAuth(_ httprouter.Handle, _ ...string) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		responder.Write401(w)
	}
}
