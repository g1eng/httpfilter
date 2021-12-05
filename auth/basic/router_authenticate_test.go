package basic

import (
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"os"
)

//this should result 200 status
func (s *authTestSuite) TestBasicAuthCredStringRouterAuthenticate(c *C) {
	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko")
	a.RouterAuthenticate(s.routeOK)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)

}

//this should result 401 status
func (s *authTestSuite) TestBasicAuthCredStringRouterUnauthorized(c *C) {
	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "pokemoko")
	a.RouterAuthenticate(s.routeOK)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusUnauthorized)

}

//this should result 400 status
func (s *authTestSuite) TestBasicAuthRouterStringBadPayload(c *C) {

	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", "Basic YnV0LWRhdGEtaXMtbm90LWhhdmUtaHRwYXNzd2QtZm9ybWF0Cg==")
	a.RouterAuthenticate(s.routeOK)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusBadRequest)

}

//this should result 200 status
func (s *authTestSuite) TestBasicAuthHtpasswdRouterAuthenticate(c *C) {
	a, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/htpasswd")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko")
	a.RouterAuthenticate(s.routeOK)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)

}
