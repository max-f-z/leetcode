package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test55(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	ans := canJump(nums)
	assert.Check(t, ans == true)
}
