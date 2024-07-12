package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2149(t *testing.T) {
	nums := []int{3, 1, -2, -5, 2, -4}
	ans := rearrangeArray(nums)
	assert.DeepEqual(t, ans, []int{3, -2, 1, -5, 2, -4})
}
