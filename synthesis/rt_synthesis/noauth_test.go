package rt_synthesis

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

func (s *rtSynthesisTestSuite) TestNoAuth(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	NoAuth(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *rtSynthesisTestSuite) TestForbid(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	Forbid(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusForbidden)
}
