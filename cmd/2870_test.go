package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2870(t *testing.T) {
	nums := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
	ans := minOperations2870(nums)
	assert.Check(t, ans == 4)

	nums = []int{2, 1, 2, 2, 3, 3}
	ans = minOperations2870(nums)
	assert.Check(t, ans == -1)

	nums = []int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}
	ans = minOperations2870(nums)
	assert.Check(t, ans == 7)
}
