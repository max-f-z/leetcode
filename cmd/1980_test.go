package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1980(t *testing.T) {
	nums := []string{"00", "10"}

	ans := findDifferentBinaryString(nums)

	assert.Check(t, ans == "01")
}
