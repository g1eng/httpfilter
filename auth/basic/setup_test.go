package basic

import (
	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
	"net/http"
	"testing"
)

func init() {
	Suite(&authTestSuite{})
}

type authTestSuite struct {
	auth *Authenticator
}

func Test(t *testing.T) { TestingT(t) }

func (s *authTestSuite) SetUpTest(_ *C) {
	//s.auth = NewBasicAuth("koremo:$apr1$7OpaItYk$P0pMHjZFCCmboF3RfrzZv.")
}

func (s *authTestSuite) routeOK(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func (s *authTestSuite) handleOK(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
