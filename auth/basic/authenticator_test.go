package basic

import (
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"os"
)

//this should result error
func (s *authTestSuite) TestNewInvalidCredString(c *C) {
	_, err := NewBasicAuth("ok")
	c.Check(err, NotNil)
}

//this should result error
func (s *authTestSuite) TestNewNoHtpasswd(c *C) {
	_, err := NewBasicAuth("/home/not/there")
	c.Check(err, NotNil)
}

//this should result error
func (s *authTestSuite) TestNewInvalidHtpasswd(c *C) {
	_, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/corrupt.htpasswd")
	c.Check(err, NotNil)
}

//this should not result error
func (s *authTestSuite) TestNewValidStringCred(c *C) {
	a, err := NewBasicAuth("koremo:$2y$12$lJD.tslxAuRtLdWamSuWcOZ4fcpwWX4VZIfj3Ph/Y9RYSKVP4NSMW")
	c.Check(err, IsNil)
	c.Check(a, NotNil)

}

//this should not result error
func (s *authTestSuite) TestNewValidHtpasswd(c *C) {
	a, err := NewBasicAuth(os.Getenv("PWD") + "/../../fixtures/htpasswd")
	c.Check(err, IsNil)
	count := 0
	for range a.userCredentials {
		count++
	}
	c.Check(count, Equals, 2)
	c.Check(a, NotNil)
}

//this should result 400 status
func (s *authTestSuite) TestBasicAuthCredStringBad(c *C) {
	a, err := NewBasicAuth("koremo:notvalid")
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

//this should result 200 status
func (s *authTestSuite) TestBasicAuthCredStringAuthenticate(c *C) {
	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko")
	a.Authenticate(s.handleOK)(w, r)
	c.Check(w.Code, Equals, http.StatusOK)

	////for httprouter
	//w = httptest.NewRecorder()
	//r = httptest.NewRequest("GET","/",nil)
	//r.SetBasicAuth("koremo","ashinokokara8")
	//a.RouterAuthenticate(s.routeOK)(w, r, nil)
	//c.Check(w.Code, Equals, http.StatusOK)
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

//this should result 200 status
func (s *authTestSuite) TestBasicAuthCredStringRouterAuthenticate(c *C) {
	a, err := NewBasicAuth("sampleuser01:mokomoko")
	c.Check(err, IsNil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.SetBasicAuth("sampleuser01", "mokomoko")
	a.RouterAuthenticate(s.routeOK)(w, r, nil)
	c.Check(w.Code, Equals, http.StatusOK)

	////for httprouter
	//w = httptest.NewRecorder()
	//r = httptest.NewRequest("GET","/",nil)
	//r.SetBasicAuth("koremo","ashinokokara8")
	//a.RouterAuthenticate(s.routeOK)(w, r, nil)
	//c.Check(w.Code, Equals, http.StatusOK)
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
