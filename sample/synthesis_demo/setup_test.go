package main

import (
	. "gopkg.in/check.v1"
	"testing"
)

func init() {
	Suite(&sampleTestSuite{})
}

type sampleTestSuite struct {
}

func Test(t *testing.T) { TestingT(t) }
