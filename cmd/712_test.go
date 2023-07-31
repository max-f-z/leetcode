package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test712(t *testing.T) {
	s1 := "sea"
	s2 := "eat"

	ans := minimumDeleteSum(s1, s2)

	assert.Check(t, ans == 231)

	s1 = "delete"
	s2 = "leet"

	ans = minimumDeleteSum(s1, s2)

	assert.Check(t, ans == 403)
}
