package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1144(t *testing.T) {
	arr := []int{9, 6, 1, 6, 2}
	ans := movesToMakeZigzag(arr)

	assert.Check(t, ans == 4)
}
