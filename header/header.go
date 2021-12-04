package header

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

// setGenericHeader is setter of CORS header for http.ResponseWriter of each RawRoute.
// This method is for common headers.
func (f *Filter) setGenericHeader(w http.ResponseWriter) {
	if f.origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", f.origin)
	}
	allowHeaders := "X-PINGOTHER,cache,Authorization"
	if f.customHeader != "" {
		allowHeaders += "," + f.customHeader
	}
	w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
	w.Header().Set("Access-Control-Max-Age", "86400")
}

// appendCorsMethodHeader is a simple setter for Access-Control-Allow-Methods header for any routes.
// This method is a wrapper of httprouter.Handle.
func (f *Filter) appendCorsMethodHeader(resource string, handler httprouter.Handle, _ ...string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		f.setGenericHeader(w)
		f.setHardeningHeader(w)
		var allowedMethods []string
		for _, method := range []string{"GET", "POST", "PUT", "DELETE", "PATCH"} {
			if h, _, _ := f.RawRoute.Lookup(method, resource); h != nil {
				allowedMethods = append(allowedMethods, method)
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ",")+",OPTIONS")
		handler(w, r, ps)
	}
}

func (f *Filter) setHardeningHeader(w http.ResponseWriter) {
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "no-sniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	//w.Header().Set("Content-Security-Policy", "default-src 'self'")
}
