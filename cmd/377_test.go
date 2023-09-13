package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test377(t *testing.T) {
	nums := []int{1, 2, 3}

	ans := combinationSum4(nums, 4)

	assert.Check(t, ans == 7)
}
