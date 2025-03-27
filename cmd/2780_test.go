package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2780(t *testing.T) {
	nums := []int{1, 2, 2, 2}
	ans := minimumIndex(nums)

	assert.Check(t, ans == 2)

	nums = []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}
	ans = minimumIndex(nums)

	assert.Check(t, ans == 4)

	nums = []int{1, 2, 1}
	ans = minimumIndex(nums)

	assert.Check(t, ans == -1)
}
