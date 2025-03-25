package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3394(t *testing.T) {
	n := 5
	rectangles := [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}

	got := checkValidCuts(n, rectangles)

	assert.Check(t, got, true)
}
