package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2366(t *testing.T) {
	nums := []int{3, 9, 3}
	ans := minimumReplacement(nums)

	assert.Check(t, ans == 2)

	nums = []int{1, 2, 3, 4, 5}
	ans = minimumReplacement(nums)

	assert.Check(t, ans == 0)

	nums = []int{2, 10, 20, 19, 1}
	ans = minimumReplacement(nums)

	assert.Check(t, ans == 47)
}
