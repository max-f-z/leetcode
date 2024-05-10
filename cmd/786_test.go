package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test786(t *testing.T) {
	arr := []int{1, 2, 3, 5}
	ans := kthSmallestPrimeFraction(arr, 3)
	assert.DeepEqual(t, ans, []int{2, 5})
}
