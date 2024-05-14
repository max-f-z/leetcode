package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1219(t *testing.T) {
	grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
	ans := getMaximumGold(grid)
	assert.Check(t, ans == 24)

	grid = [][]int{{0, 0, 34, 0, 5, 0, 7, 0, 0, 0}, {0, 0, 0, 0, 21, 0, 0, 0, 0, 0}, {0, 18, 0, 0, 8, 0, 0, 0, 4, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {15, 0, 0, 0, 0, 22, 0, 0, 0, 21}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 7, 0, 0, 0, 0, 0, 0, 38, 0}}
	ans = getMaximumGold(grid)
	assert.Check(t, ans == 38)
}
