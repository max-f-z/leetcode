package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2849(t *testing.T) {
	ans := isReachableAtTime(2, 4, 7, 7, 6)
	assert.Check(t, ans)
}
