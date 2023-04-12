package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test71(t *testing.T) {
	a := "/home/../../asdf"

	ans := simplifyPath(a)

	assert.Check(t, ans == "/home/asdf")
}
