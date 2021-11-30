package core

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http/httptest"
)

func (s *filterTestSuite) TestGlobalHeader(c *C) {
	checkedHeaders := []string{
		"Access-Control-Allow-Headers",
		"Access-Control-Max-Age",
	}
	w := *httptest.NewRecorder()
	for _, v := range checkedHeaders {
		c.Check(w.Header().Get(v), Equals, "")
	}
	s.f.setGenericHeader(&w)
	for _, v := range checkedHeaders {
		c.Check(w.Header().Get(v), Not(Equals), "")
	}
	//but cors origin header is blank string by default
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, "")
}

func (s *filterTestSuite) TestGenericHeader(c *C) {
	w := *httptest.NewRecorder()
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, "")
	s.f.origin = "this"
	s.f.setGenericHeader(&w)
	//cors origin header is set by origin field
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, s.f.origin)

	// reset
	w = *httptest.NewRecorder()
	s.f.origin = ""
	s.f.setGenericHeader(&w)
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, "")
	// and cors origin header is set by setOrigin
	s.f.SetOrigin("that")
	s.f.setGenericHeader(&w)
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, s.f.origin)
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, "that")
}

func (s *filterTestSuite) TestGenericCustomHeader(c *C) {
	w := *httptest.NewRecorder()
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, "")
	s.f.customHeader = "X_CSRF_DUMMY_TOKEN"
	s.f.setGenericHeader(&w)
	c.Check(w.Header().Get("Access-Control-Allow-Headers"), Matches, "^.+,"+s.f.customHeader+"$")
}

//TestAllow is wrapper case for setGenericHeader and setHardeningHeader
func (s *filterTestSuite) TestAllow(c *C) {
	s.f.SetOrigin("that")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ok", bytes.NewBufferString("this is ok"))
	c.Check(w.Header().Get("Access-Control-Allow-Methods"), Equals, "")
	s.f.appendCorsMethodHeader("/ok", s.nullHandler)(w, r, nil)
	c.Check(w.Header().Get("X-Content-Type-Options"), Equals, "no-sniff")
}
