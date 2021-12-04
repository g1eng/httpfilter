package rt_synthesis

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

func (s *filterTestSuite) TestNoAuth(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	NoAuth(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestForbid(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	Forbid(s.echoResponder)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusForbidden)
}
