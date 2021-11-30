package ipfilter

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
)

//IPHandler is (experimental) one of the AuthWrapper which enables IP filtering
//to allow/deny specific network address.
//FIXME: matching for IP section is only supported at now and the subnet is ignored
func (ipf IPFilter) IPHandler(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		remoteIP := net.ParseIP(r.RemoteAddr)
		if ipf.Policy && ipf.IPAddr.Equal(remoteIP) {
			handle(w, r, ps)
		} else if !ipf.Policy && !ipf.IPAddr.Equal(remoteIP) {
			handle(w, r, ps)
		} else {
			return
		}
	}
}
