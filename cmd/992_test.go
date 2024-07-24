package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test992(t *testing.T) {
	nums := []int{1, 2, 1, 2, 3}
	ans := subarraysWithKDistinct(nums, 2)
	assert.Check(t, ans == 7)
}
