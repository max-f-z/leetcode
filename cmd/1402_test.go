package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1402(t *testing.T) {
	satisfaction := []int{-1, -8, 0, 5, -7}

	ans := maxSatisfaction(satisfaction)

	assert.Check(t, ans == 14)
}
