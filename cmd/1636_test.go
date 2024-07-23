package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1636(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 3}
	ans := frequencySort(nums)
	assert.DeepEqual(t, ans, []int{3, 1, 1, 2, 2, 2})

	nums = []int{2, 3, 1, 3, 2}
	ans = frequencySort(nums)
	assert.DeepEqual(t, ans, []int{1, 3, 3, 2, 2})
}
