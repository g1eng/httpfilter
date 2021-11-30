package core

import (
	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func init() {
	Suite(&filterTestSuite{})
}

type filterTestSuite struct {
	f      HttpFilter
	dummy  HttpFilter
	result routeResult
}
type routeResult struct {
	writer  http.ResponseWriter
	request *http.Request
}

func Test(t *testing.T) { TestingT(t) }

func (s *filterTestSuite) SetUpTest(_ *C) {
	s.f = *NewFilter()
	s.dummy = *NewFilter()
	s.result = routeResult{}
}

func noAuth(handle httprouter.Handle) httprouter.Handle {
	return handle
}

func (s *filterTestSuite) nullHandler(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	return
}

//echoResponder is a test stub for POST requests, which makes echo of a post body
func (s *filterTestSuite) echoResponder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var a []byte
	_, _ = r.Body.Read(a)
	_, _ = w.Write(a)
}
