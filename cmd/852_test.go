package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test852(t *testing.T) {
	arr := []int{0, 1, 0}
	ans := peakIndexInMountainArray(arr)

	assert.Check(t, ans == 1)
}
