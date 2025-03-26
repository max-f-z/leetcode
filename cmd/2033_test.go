package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2033(t *testing.T) {
	grid := [][]int{{2, 4}, {6, 8}}
	x := 2

	got := minOperations2033(grid, x)

	assert.Check(t, got == 4)
}
