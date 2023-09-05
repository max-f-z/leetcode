package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2845(t *testing.T) {
	nums := []int{3, 2, 4}
	ans := countInterestingSubarrays(nums, 2, 1)
	assert.Check(t, ans == 3)

	nums = []int{3, 1, 9, 6}
	ans = countInterestingSubarrays(nums, 3, 0)
	assert.Check(t, ans == 2)
}
