package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2551(t *testing.T) {
	weights := []int{1, 3, 5, 1}
	k := 2

	ans := putMarbles(weights, k)
	assert.Check(t, ans == 4)
}
