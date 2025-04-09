package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3375(t *testing.T) {
	nums := []int{9, 7, 5, 3}
	k := 1
	ans := minOperations3375(nums, k)
	assert.Check(t, ans == 4)
}
