package main

import (
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
)

func (s *filterTestSuite) TestOha(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8080/o/ha", nil)
	oHandler(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestConch(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8080/o/con", nil)
	conchHandler(w, r)
	c.Check(w.Code, Equals, http.StatusAccepted)
}
