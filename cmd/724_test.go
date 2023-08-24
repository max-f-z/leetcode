package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test724(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}

	ans := pivotIndex(nums)

	assert.Check(t, ans == 3)
}
