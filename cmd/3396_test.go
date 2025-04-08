package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3396(t *testing.T) {
	nums := []int{1, 2, 3, 4, 2, 3, 3, 5, 7}
	ans := minimumOperations3396(nums)
	assert.Check(t, ans == 2)
}
