package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test131(t *testing.T) {
	n := 6
	connections := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}

	ans := makeConnected(n, connections)

	assert.Check(t, ans == 2)
}
