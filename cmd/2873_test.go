package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2873(t *testing.T) {
	nums := []int{12, 6, 1, 2, 7}
	ans := maximumTripletValue(nums)
	assert.Check(t, ans == 77)

	nums = []int{10, 13, 6, 2}
	ans = maximumTripletValue(nums)
	assert.Check(t, ans == 14)
}
