package header

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

func (f *Filter) GET(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.GET(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *Filter) POST(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.POST(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *Filter) PUT(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.PUT(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *Filter) DELETE(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.DELETE(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *Filter) PATCH(resource string, authenticator rt_synthesis.AuthWrapper, handler httprouter.Handle) {
	f.RawRoute.PATCH(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(authenticator(handler))))
}

func (f *Filter) OPTIONS(resource string, _ httprouter.Handle) {
	f.RawRoute.OPTIONS(resource, f.appendCorsMethodHeader(resource, f.defaultAuth(f.corsRequestHandler)))
}

func (f *Filter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.RawRoute.ServeHTTP(w, r)
}

func (f *Filter) Route() *httprouter.Router {
	return f.RawRoute
}
