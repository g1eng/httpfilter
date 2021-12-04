package synthesis

import (
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func init() {
	Suite(&filterTestSuite{})
}

type filterTestSuite struct {
}

func Test(t *testing.T) { TestingT(t) }

func (s *filterTestSuite) SetUpTest(_ *C) {
}

//plainEchoResponder is a test stub for POST requests, which makes echo of a post body
func (s *filterTestSuite) echoResponder(w http.ResponseWriter, r *http.Request) {
	var a []byte
	_, _ = r.Body.Read(a)
	_, _ = w.Write(a)
}
