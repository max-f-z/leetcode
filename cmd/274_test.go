package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test274(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	ans := hIndex(citations)
	assert.Check(t, ans == 3)

	citations = []int{1}
	ans = hIndex(citations)
	assert.Check(t, ans == 1)
}
