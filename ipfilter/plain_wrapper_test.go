package ipfilter

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

func (s *filterTestSuite) TestPlainAllowedIP(c *C) {
	w := httptest.NewRecorder()
	//allow from 127.0.0.1
	s.f = NewIPFilter(true, []string{"127.0.0.1"})
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	s.f.PlainIPHandler(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestPlainDeniedIP(c *C) {
	w := httptest.NewRecorder()
	//allow from 127.0.0.1
	s.f = NewIPFilter(false, []string{"127.0.0.1"})
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	s.f.PlainIPHandler(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}

func (s *filterTestSuite) TestPlainAllowedSubnet(c *C) {
	w := httptest.NewRecorder()
	//allow from 127.0.0.1
	s.f = NewIPFilter(true, []string{"127.0.0.0/8"})
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	s.f.PlainIPHandler(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestPlainDeniedSubnet(c *C) {
	w := httptest.NewRecorder()
	//allow from 127.0.0.1
	s.f = NewIPFilter(false, []string{"127.0.0.0/8"})
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString(""))
	s.f.PlainIPHandler(s.plainEchoResponder)(w, r)
	c.Check(w.Code, Equals, http.StatusForbidden)
}
