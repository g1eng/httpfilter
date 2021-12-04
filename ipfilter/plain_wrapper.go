package ipfilter

import (
	"github.com/g1eng/httpfilter/session/responder"
	"net"
	"net/http"
)

//PlainIPHandler is one of the AuthWrapper which enables IP filtering
//to allow/deny specific network address.
func (ipf *IPFilter) PlainIPHandler(handle http.HandlerFunc, _ ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
		if ipf.RawFilter.Allowed(remoteIP) {
			handle(w, r)
		} else {
			responder.Write403(w)
		}
	}
}
