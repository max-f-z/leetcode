package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test128(t *testing.T) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	ans := longestConsecutiveII(nums)

	assert.Check(t, ans == 9)
}
