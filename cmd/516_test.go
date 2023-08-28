package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test516(t *testing.T) {
	str := "bbbab"

	ans := longestPalindromeSubseq(str)

	assert.Check(t, ans == 4)
}
