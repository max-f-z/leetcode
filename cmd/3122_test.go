package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3122(t *testing.T) {
	grid := [][]int{{2, 9}, {3, 7}}
	ans := minimumOperations3122(grid)
	assert.Check(t, ans == 2)

	grid = [][]int{{1, 1, 1}, {0, 0, 0}}
	ans = minimumOperations3122(grid)
	assert.Check(t, ans == 3)

}
