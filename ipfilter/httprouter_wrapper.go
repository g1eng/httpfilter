package ipfilter

import (
	"github.com/g1eng/httpfilter/session/responder"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
)

//IPHandler is (experimental) one of the AuthWrapper which enables IP filtering
//to allow/deny specific network address.
func (ipf *IPFilter) IPHandler(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
		if ipf.RawFilter.Allowed(remoteIP) {
			handle(w, r, ps)
		} else {
			responder.Write403(w)
		}
	}
}
