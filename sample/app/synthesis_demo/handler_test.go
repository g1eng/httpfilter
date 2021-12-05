package main

import (
	"github.com/julienschmidt/httprouter"
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

func (s *filterTestSuite) TestConchRoute400(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8080/o/con", nil)
	ps := httprouter.Params{
		httprouter.Param{Key: "nok", Value: "tora"},
	}
	conchRoute(w, r, ps)
	c.Check(w.Code, Equals, http.StatusBadRequest)
}

func (s *filterTestSuite) TestConchRoute202(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8080/o/con", nil)
	ps := httprouter.Params{
		httprouter.Param{Key: "ok", Value: "torari"},
	}
	conchRoute(w, r, ps)
	c.Check(w.Code, Equals, http.StatusAccepted)
}

func (s *filterTestSuite) TestORoute202(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8080/o/ha", nil)
	ps := httprouter.Params{
		httprouter.Param{Key: "ok", Value: "torari"},
	}
	conchRoute(w, r, ps)
	c.Check(w.Code, Equals, http.StatusAccepted)
}

func (s *filterTestSuite) TestORoute400(c *C) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8080/o/ha", nil)
	ps := httprouter.Params{
		httprouter.Param{Key: "nok", Value: "tora"},
	}
	conchRoute(w, r, ps)
	c.Check(w.Code, Equals, http.StatusBadRequest)
}
