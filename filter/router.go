package filter

import (
	"github.com/g1eng/httpfilter/synthesis/rt_synthesis"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HttpHandlerHandler interface {
	GET(string, rt_synthesis.AuthWrapper, httprouter.Handle)
	POST(string, rt_synthesis.AuthWrapper, httprouter.Handle)
	PUT(string, rt_synthesis.AuthWrapper, httprouter.Handle)
	DELETE(string, rt_synthesis.AuthWrapper, httprouter.Handle)
	PATCH(string, rt_synthesis.AuthWrapper, httprouter.Handle)
	OPTIONS(string, httprouter.Handle)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func (f *HttpFilter) GET(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.GET(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *HttpFilter) POST(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.POST(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *HttpFilter) PUT(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.PUT(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *HttpFilter) DELETE(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.DELETE(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *HttpFilter) PATCH(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.PATCH(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *HttpFilter) OPTIONS(resource string, _ httprouter.Handle) {
	f.RawRoute.OPTIONS(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(f.corsRequestHandler)))
}

func (f *HttpFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.RawRoute.ServeHTTP(w, r)
}

func (f *HttpFilter) Route() *httprouter.Router {
	return f.RawRoute
}
