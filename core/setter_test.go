package core

import (
	. "gopkg.in/check.v1"
)

func (s *filterTestSuite) TestSetDebug(c *C) {
	c.Check(s.f.debug, Equals, false)
	s.f.SetDebug()
	c.Check(s.f.debug, Equals, true)
}

func (s *filterTestSuite) TestSetOrigin(c *C) {
	hostname := "http://some.example.org"
	c.Check(s.f.origin, Equals, "")
	s.f.SetOrigin(hostname)
	c.Check(s.f.origin, Equals, hostname)
}

func (s *filterTestSuite) TestSetCustomHeader(c *C) {
	customHeader := "X_CSRF_DUMMY_TOKEN"
	c.Check(s.f.customHeader, Equals, "")
	s.f.SetCustomHeader(customHeader)
	c.Check(s.f.customHeader, Equals, customHeader)
}
