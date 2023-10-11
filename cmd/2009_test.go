package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2009(t *testing.T) {
	nums := []int{4, 2, 5, 3}
	ans := minOperations(nums)
	assert.Check(t, ans == 0)

	nums = []int{1, 2, 3, 5, 6}
	ans = minOperations(nums)
	assert.Check(t, ans == 1)

	nums = []int{1, 10, 100, 1000}
	ans = minOperations(nums)
	assert.Check(t, ans == 3)
}
