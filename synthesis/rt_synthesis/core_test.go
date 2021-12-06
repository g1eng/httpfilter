package rt_synthesis

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

//200
func (s *rtSynthesisTestSuite) TestAndOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(NoAuth, NoAuth)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *rtSynthesisTestSuite) TestAndOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(NoAuth, Forbid)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *rtSynthesisTestSuite) TestAndForbiddenOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(Forbid, NoAuth)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//403
func (s *rtSynthesisTestSuite) TestAndForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAND(Forbid, Forbid)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//200
func (s *rtSynthesisTestSuite) TestOrOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(NoAuth, NoAuth)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)
}

//200
func (s *rtSynthesisTestSuite) TestOrOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(NoAuth, Forbid)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *rtSynthesisTestSuite) TestOrForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthOR(Forbid, Forbid)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *rtSynthesisTestSuite) TestAuthAllOK3(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *rtSynthesisTestSuite) TestAuthAllOK5(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth, NoAuth, NoAuth)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *rtSynthesisTestSuite) TestAuthAllOK4Not1(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	AuthAll(NoAuth, NoAuth, NoAuth, Forbid, NoAuth)(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusForbidden)
}
