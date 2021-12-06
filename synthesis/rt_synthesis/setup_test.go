package rt_synthesis

import (
	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func init() {
	Suite(&rtSynthesisTestSuite{})
}

type rtSynthesisTestSuite struct {
}

func Test(t *testing.T) { TestingT(t) }

func (s *rtSynthesisTestSuite) SetUpTest(_ *C) {
}

func (s *rtSynthesisTestSuite) nullHandler(_ http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	return
}

//echoResponder is a test stub for POST requests, which makes echo of a post body
func (s *rtSynthesisTestSuite) echoResponder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var a []byte
	_, _ = r.Body.Read(a)
	_, _ = w.Write(a)
}

//plainEchoResponder is a test stub for POST requests, which makes echo of a post body
func (s *rtSynthesisTestSuite) plainEchoResponder(w http.ResponseWriter, r *http.Request) {
	var a []byte
	_, _ = r.Body.Read(a)
	_, _ = w.Write(a)
}
