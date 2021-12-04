package basic

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
	auth *Authenticator
}

func Test(t *testing.T) { TestingT(t) }

func (s *filterTestSuite) SetUpTest(_ *C) {
	//s.auth = NewBasicAuth("koremo:$apr1$7OpaItYk$P0pMHjZFCCmboF3RfrzZv.")
}

func (s *filterTestSuite) routeOK(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func (s *filterTestSuite) handleOK(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
