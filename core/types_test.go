package core

import (
	. "gopkg.in/check.v1"
)

func (s *filterTestSuite) TestNewFilter(c *C) {
	f := *NewFilter(nil)
	c.Check(f.origin, Equals, "")
	c.Check(f.origin, Equals, s.dummy.origin)
	c.Check(f.debug, Equals, false)
	c.Check(f.debug, Equals, s.dummy.debug)
	c.Log("It must set httprouter.Router instance and a authorization handler")
	c.Check(f.RawRoute, NotNil)
	c.Check(f.defaultAuth, NotNil)
}
