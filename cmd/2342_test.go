package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2342(t *testing.T) {
	nums := []int{18, 43, 36, 13, 7}
	ans := maximumSum(nums)

	assert.Check(t, ans == 54)
}
