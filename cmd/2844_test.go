package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2844(t *testing.T) {
	num := "2245047"
	ans := minimumOperations(num)
	assert.Check(t, ans == 2)

	num = "2908305"
	ans = minimumOperations(num)
	assert.Check(t, ans == 3)

	num = "10"
	ans = minimumOperations(num)
	assert.Check(t, ans == 1)
}
