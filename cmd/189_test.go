package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test189(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate189(nums, k)
	assert.DeepEqual(t, nums, []int{5, 6, 7, 1, 2, 3, 4})
}
