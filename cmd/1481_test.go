package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1481(t *testing.T) {
	arr := []int{5, 5, 4}
	ans := findLeastNumOfUniqueInts(arr, 1)
	assert.Check(t, ans == 1)

	arr = []int{4, 3, 1, 1, 3, 3, 2}
	ans = findLeastNumOfUniqueInts(arr, 3)
	assert.Check(t, ans == 2)
}
