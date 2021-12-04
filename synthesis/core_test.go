package synthesis

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

//200
func (s *filterTestSuite) TestAndOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *filterTestSuite) TestAndOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(NoAuth, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *filterTestSuite) TestAndForbiddenOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(Forbid, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//403
func (s *filterTestSuite) TestAndForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(Forbid, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//200
func (s *filterTestSuite) TestOrOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//200
func (s *filterTestSuite) TestOrOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(NoAuth, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *filterTestSuite) TestOrForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(Forbid, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *filterTestSuite) TestAuthAllOK3(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestAuthAllOK5(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth, NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestAuthAllOK4Forbidden1(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth, Forbid, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}
