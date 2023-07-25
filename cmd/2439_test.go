package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2439(t *testing.T) {
	nums := []int{4, 100, 1, 6}
	ans := minimizeArrayValue(nums)

	assert.Check(t, ans == 52)
}
