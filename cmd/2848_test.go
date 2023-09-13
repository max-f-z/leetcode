package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2848(t *testing.T) {
	nums := [][]int{{3, 6}, {1, 5}, {4, 7}}
	ans := numberOfPoints(nums)

	assert.Check(t, ans == 7)
}
