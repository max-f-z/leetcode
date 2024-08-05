package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2134(t *testing.T) {
	nums := []int{0, 1, 0, 1, 1, 0, 0}
	ans := minSwaps(nums)
	assert.Check(t, ans == 1)

	nums = []int{0, 1, 1, 1, 0, 0, 1, 1, 0}
	ans = minSwaps(nums)
	assert.Check(t, ans == 2)
}
