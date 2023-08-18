package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2815(t *testing.T) {
	nums := []int{51, 71, 17, 24, 42}

	ans := maxSum(nums)

	assert.Check(t, ans == 88)

	nums = []int{84, 91, 18, 59, 27, 9, 81, 33, 17, 58}

	ans = maxSum(nums)

	assert.Check(t, ans == 165)
}
