package basic

import (
	. "gopkg.in/check.v1"
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
