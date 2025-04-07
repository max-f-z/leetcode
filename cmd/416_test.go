package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test416(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	ans := canPartition(nums)
	assert.Check(t, ans == true)

	nums = []int{2, 1, 5}
	ans = canPartition(nums)
	assert.Check(t, ans == false)

	nums = []int{100, 100, 99, 97}
	ans = canPartition(nums)
	assert.Check(t, ans == false)
}
