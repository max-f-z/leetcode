package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1910(t *testing.T) {
	s := "daabcbaabcbc"
	part := "abc"

	ans := removeOccurrences(s, part)
	assert.Check(t, ans == "dab")
}
