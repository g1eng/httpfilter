package core

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// corsRequestHandler is the handler for OPTION request, that simply returns CORS header allowed for the resource.
// At now, this is globally applied to OPTIONS requests in /*path.
// This is thought to be insecure for reasons.
func (f *HttpFilter) corsRequestHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	f.setGenericHeader(w)
	f.setHardeningHeader(w)
	w.WriteHeader(http.StatusOK)
}

func NullHandler(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	return
}
