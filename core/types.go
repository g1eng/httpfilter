package core

import (
	"github.com/g1eng/w3fs/middleware/filter/wrapper"
	"github.com/julienschmidt/httprouter"
)

// HttpFilter provides filter mechanism based on wrapper.AuthWrapper or response headers related to
// CORS or other security mechanisms. It acts as an overlay of httprouter.Router.
type HttpFilter struct {
	origin       string              //cors origin
	customHeader string              //customHeader available on CORS request
	debug        bool                //debug flag for HttpFilter
	RawRoute     *httprouter.Router  //RawRoute field holds raw httprouter.Router multiplexer
	defaultAuth  wrapper.AuthWrapper //defaultAuth is the default auth wrapper for any routes. It is used to generate authorized RawRoute.
	route        map[string][]string //route map for acceptable method sets (GET, POST, PUT, etc.)
}

// NewFilter generate new http filter instance associated to httprouter.Router instance.
// An authorizer function can be set to authorize/authenticate users.
// If nil authorizer given, authorization process for routes will be removed.
func NewFilter(authorizer ...wrapper.AuthWrapper) *HttpFilter {
	var auth wrapper.AuthWrapper
	if authorizer == nil || authorizer[0] == nil {
		auth = wrapper.AuthAll(authorizer...)
	} else {
		auth = wrapper.NoAuth
	}
	return &HttpFilter{
		origin:       "",
		customHeader: "",
		debug:        false,
		RawRoute:     httprouter.New(),
		defaultAuth:  auth,
		route:        map[string][]string{},
	}
}
