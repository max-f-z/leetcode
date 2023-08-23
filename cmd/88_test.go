package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test88(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	mergeSorted(nums1, m, nums2, n)

	assert.Check(t, len(nums1) == 6)
}
