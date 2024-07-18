package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test268(t *testing.T) {
	nums := []int{3, 0, 1}
	ans := missingNumber(nums)

	assert.Check(t, ans == 2)
}
