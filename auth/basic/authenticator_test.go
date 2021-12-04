package basic

import (
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"os"
)

//this should result error
func (s *filterTestSuite) TestNewInvalidCredString(c *C) {
	_, err := NewBasicAuth("ok")
	c.Check(err, NotNil)
}

//this should result error
func (s *filterTestSuite) TestNewNoHtpasswd(c *C) {
	_, err := NewBasicAuth("/home/not/there")
	c.Check(err, NotNil)
}

//this should result error
func (s *filterTestSuite) TestNewInvalidHtpasswd(c *C) {
	_, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/corrupt.htpasswd")
	c.Check(err, NotNil)
}

//this should not result error
func (s *filterTestSuite) TestNewValidStringCred(c *C) {
	a, err := NewBasicAuth("koremo:$apr1$DrWuZAEw$pwnhPomgEICGtAy1qZWWY0")
	c.Check(err, IsNil)
	c.Check(a, NotNil)

}

//this should not result error
func (s *filterTestSuite) TestNewValidHtpasswd(c *C) {
	a, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/htpasswd")
	c.Check(err, IsNil)
	c.Check(a, NotNil)
}

func (s *filterTestSuite) TestBasicAuthCredStringForbidden(c *C) {
	a, err := NewBasicAuth("koremo:$apr1$DrWuZAEw$pwnhPomgEICGtAy1qZWWY0")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusBadRequest)

	////for httprouter
	//w = httptest.NewRecorder()
	//r = httptest.NewRequest("GET","/",nil)
	//a.RouterAuthenticate(s.routeOK)(w, r, nil)
	//c.Check(w.Code, Equals, http.StatusUnauthorized)
}

func (s *filterTestSuite) TestBasicAuthCredStringAuthenticate(c *C) {
	a, err := NewBasicAuth("koremo:$apr1$DrWuZAEw$pwnhPomgEICGtAy1qZWWY0")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("koremo", "ashinokokara8")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)

	////for httprouter
	//w = httptest.NewRecorder()
	//r = httptest.NewRequest("GET","/",nil)
	//r.SetBasicAuth("koremo","ashinokokara8")
	//a.RouterAuthenticate(s.routeOK)(w, r, nil)
	//c.Check(w.Code, Equals, http.StatusOK)
}

func (s *filterTestSuite) TestBasicAuthHtpasswdAuthenticate(c *C) {
	a, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/htpasswd")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)

}
