package synthesis

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

//200
func (s *synthesisTestSuite) TestAndOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *synthesisTestSuite) TestAndOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(NoAuth, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *synthesisTestSuite) TestAndForbiddenOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(Forbid, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//403
func (s *synthesisTestSuite) TestAndForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(Forbid, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//200
func (s *synthesisTestSuite) TestOrOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//200
func (s *synthesisTestSuite) TestOrOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(NoAuth, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *synthesisTestSuite) TestOrForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(Forbid, Forbid)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *synthesisTestSuite) TestAuthAllOK3(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *synthesisTestSuite) TestAuthAllOK5(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth, NoAuth, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *synthesisTestSuite) TestAuthAllOK4Forbidden1(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth, Forbid, NoAuth)(s.echoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}
