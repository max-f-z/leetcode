package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test169(t *testing.T) {
	nums := []int{3, 2, 3}
	ans := majorityElement(nums)
	assert.Check(t, ans == 3)
}
