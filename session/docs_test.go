package session

import (
	. "gopkg.in/check.v1"
	"testing"
)

func init() {
	Suite(&filterTestSuite{})
}

type filterTestSuite struct {
}

func Test(t *testing.T) { TestingT(t) }

func (s *filterTestSuite) SetUpTest(_ *C) {
}
