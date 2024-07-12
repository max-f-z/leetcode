package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1717(t *testing.T) {
	s := "cdbcbbaaabab"
	ans := maximumGain(s, 4, 5)
	assert.Check(t, ans == 19)

	s = "aabbaaxybbaabb"
	ans = maximumGain(s, 5, 4)
	assert.Check(t, ans == 20)
}
