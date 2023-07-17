package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1135(t *testing.T) {
	relations := [][]int{{1, 2}, {2, 3}, {3, 1}}

	ans := minimumSemesters(3, relations)

	assert.Check(t, ans == -1)
}
