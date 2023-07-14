package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1218(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	ans := longestSubsequence(arr, 1)
	assert.Check(t, ans == 4)

	arr = []int{1, 3, 5, 7, 9}
	ans = longestSubsequence(arr, 1)
	assert.Check(t, ans == 1)

	arr = []int{1, 3, 5, 7, 9}
	ans = longestSubsequence(arr, 2)
	assert.Check(t, ans == 5)
}
