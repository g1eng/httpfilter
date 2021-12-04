package header

import (
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"strings"
)

// integration test for GET resources
func (s *filterTestSuite) TestGet(c *C) {
	target := "/get/ok"
	hostname := "http://example.com"

	s.f.SetOrigin(hostname)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", target, nil)
	s.f.GET(target, noAuth, s.nullHandler)
	h, p, _ := s.f.RawRoute.Lookup("GET", target)
	h(w, r, p)
	c.Check(w.Code, Equals, http.StatusOK)
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, hostname)
	c.Check(w.Header().Get("Access-Control-Allow-Methods"), Equals, "GET,OPTIONS")
	c.Check(w.Header().Get("X-Content-Type-Options"), Equals, "no-sniff")
}

// integration test for POST resources
func (s *filterTestSuite) TestPost(c *C) {
	hostname := "http://example.com"
	s.f.SetOrigin(hostname)
	w := httptest.NewRecorder()
	msg := "this is posted"
	body := strings.NewReader(msg)
	r, _ := http.NewRequest("POST", "/post/ok", body)
	s.f.POST("/post/ok", noAuth, s.echoResponder)
	h, p, _ := s.f.RawRoute.Lookup("POST", "/post/ok")
	h(w, r, p)

	c.Check(w.Code, Equals, http.StatusOK)
	c.Check(w.Header().Get("Access-Control-Allow-Origin"), Equals, hostname)
	c.Check(w.Header().Get("Access-Control-Allow-Methods"), Equals, "POST,OPTIONS")
	c.Check(w.Header().Get("X-Content-Type-Options"), Equals, "no-sniff")

	//FIXME: test body payload
	//c.Check(w.Body.String(), Equals, msg)
}

func (s *filterTestSuite) TestGetPost(c *C) {
	w := httptest.NewRecorder()
	body := strings.NewReader("this is posted")
	r, _ := http.NewRequest("GET", "/post/ok", body)

	s.f.GET("/post/ok", noAuth, s.echoResponder)
	s.f.POST("/post/ok", noAuth, s.echoResponder)
	h, p, _ := s.f.RawRoute.Lookup("GET", "/post/ok")
	h(w, r, p)
	c.Check(w.Code, Equals, http.StatusOK)
	c.Check(w.Header().Get("Access-Control-Allow-Methods"), Equals, "GET,POST,OPTIONS")
}

func (s *filterTestSuite) TestPostPutDelete(c *C) {
	w := httptest.NewRecorder()
	body := strings.NewReader("this is posted")
	r, _ := http.NewRequest("GET", "/post/ok", body)

	s.f.POST("/post/ok", noAuth, s.echoResponder)
	s.f.DELETE("/post/ok", noAuth, s.echoResponder)
	s.f.PUT("/post/ok", noAuth, s.echoResponder)
	h, p, _ := s.f.RawRoute.Lookup("POST", "/post/ok")
	h(w, r, p)
	c.Check(w.Code, Equals, http.StatusOK)
	c.Check(w.Header().Get("Access-Control-Allow-Methods"), Equals, "POST,PUT,DELETE,OPTIONS")
}
