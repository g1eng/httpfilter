package synthesis

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

//200
func (s *filterTestSuite) TestPlainAndOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthAND(PlainNoAuth, PlainNoAuth)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *filterTestSuite) TestPlainAndOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthAND(PlainNoAuth, PlainFalse)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *filterTestSuite) TestPlainAndForbiddenOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthAND(PlainFalse, PlainNoAuth)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//403
func (s *filterTestSuite) TestPlainAndForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthAND(PlainFalse, PlainFalse)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

//200
func (s *filterTestSuite) TestPlainOrOKOK(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthOR(PlainNoAuth, PlainNoAuth)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//200
func (s *filterTestSuite) TestPlainOrOKForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthOR(PlainNoAuth, PlainFalse)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//403
func (s *filterTestSuite) TestPlainOrForbiddenForbidden(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthOR(PlainFalse, PlainFalse)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *filterTestSuite) TestPlainPlainAuthAllOK3(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthAll(PlainNoAuth, PlainNoAuth, PlainNoAuth)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestPlainPlainAuthAllOK5(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthAll(PlainNoAuth, PlainNoAuth, PlainNoAuth, PlainNoAuth, PlainNoAuth)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestPlainPlainAuthAllOK4Not1(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	PlainAuthAll(PlainNoAuth, PlainNoAuth, PlainNoAuth, PlainFalse, PlainNoAuth)(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}
