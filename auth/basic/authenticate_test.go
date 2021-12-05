package basic

import (
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"os"
)

//this should result 400 status
func (s *authTestSuite) TestBasicAuthCredStringBad(c *C) {
	a, err := NewBasicAuth("koremo:notvalid")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusBadRequest)

}

//this should result 200 status
func (s *authTestSuite) TestBasicAuthCredStringAuthenticate(c *C) {
	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)
}

//this should result 200 status
func (s *authTestSuite) TestBasicAuthHtpasswdAuthenticate(c *C) {
	a, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/htpasswd")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)

}

//this should result 401 status
func (s *authTestSuite) TestBasicAuthHtpasswdUnauthorized(c *C) {
	a, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/htpasswd")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko2")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusUnauthorized)

}

//this should result 400 status
func (s *authTestSuite) TestBasicAuthStringBadHeader(c *C) {

	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", "Basic BUT-header-is-not-base64-encoded")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusBadRequest)

}

//this should result 400 status
func (s *authTestSuite) TestBasicAuthStringBadPayload(c *C) {

	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", "Basic YnV0LWRhdGEtaXMtbm90LWhhdmUtaHRwYXNzd2QtZm9ybWF0Cg==")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusBadRequest)

}

//this should result 200 status
func (s *authTestSuite) TestBasicAuthHtpasswdAuthenticate2(c *C) {
	a, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/htpasswd")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser02", "mokomoko2")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)

}
