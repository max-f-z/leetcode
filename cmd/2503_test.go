package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2503(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{2, 5, 7},
		{3, 5, 1},
	}
	queries := []int{5, 6, 2}

	ans := maxPoints(grid, queries)

	assert.Check(t, ans[0] == 5)
	assert.Check(t, ans[1] == 8)
	assert.Check(t, ans[2] == 1)
}
