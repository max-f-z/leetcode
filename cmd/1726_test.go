package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1726(t *testing.T) {
	nums := []int{2, 3, 4, 6}
	ans := tupleSameProduct(nums)

	assert.Check(t, ans == 8)
}
