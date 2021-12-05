package ipfilter

import (
	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func init() {
	Suite(&ipFilterTestSuite{})
}

type ipFilterTestSuite struct {
	f     *IPFilter
	dummy *IPFilter
}

func Test(t *testing.T) { TestingT(t) }

func (s *ipFilterTestSuite) SetUpTest(_ *C) {
}

func noAuth(handle httprouter.Handle, _ ...string) httprouter.Handle {
	return handle
}

func (s *ipFilterTestSuite) nullHandler(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	return
}

//echoResponder is a test stub for POST requests, which makes echo of a post body
func (s *ipFilterTestSuite) echoResponder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var a []byte
	_, _ = r.Body.Read(a)
	_, _ = w.Write(a)
}

func (s *ipFilterTestSuite) plainEchoResponder(w http.ResponseWriter, r *http.Request) {
	var a []byte
	_, _ = r.Body.Read(a)
	_, _ = w.Write(a)
}
