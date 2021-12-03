package filter

import (
	"bytes"
	. "gopkg.in/check.v1"
	"net/http/httptest"
)

func (s *filterTestSuite) TestHardeningHeader(c *C) {
	w := httptest.NewRecorder()
	c.Check(w.Header().Get("X-Frame-Options"), Equals, "")
	c.Check(w.Header().Get("X-Content-Type-Options"), Equals, "")
	c.Check(w.Header().Get("X-XSS-Protection"), Equals, "")
	c.Check(w.Header().Get("Content-Security-Policy"), Equals, "")

	s.f.setHardeningHeader(w)

	c.Check(w.Header().Get("X-Frame-Options"), Equals, "DENY")
	c.Check(w.Header().Get("X-Content-Type-Options"), Equals, "no-sniff")
	c.Check(w.Header().Get("X-XSS-Protection"), Equals, "1; mode=block")
	//c.Check(w.Header().Get("Content-Security-Policy"), Equals, "default-src 'self'")
}

func (s *filterTestSuite) TestCorsRequestHandler(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("OPTIONS", "/some/path", bytes.NewBufferString(""))
	s.f.origin = "http://hoge.exmaple.com"
	s.f.corsRequestHandler(w, r, nil)
	c.Check(w.Code, Equals, 200)
	c.Check(w.Header().Get("X-Content-Type-Options"), Equals, "no-sniff")
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Not(Equals), "")
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, s.f.origin)
}
