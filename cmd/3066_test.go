package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3066(t *testing.T) {
	nums := []int{2, 11, 10, 1, 3}
	k := 10

	ans := minOperations3066(nums, k)
	assert.Check(t, ans == 2)
}
