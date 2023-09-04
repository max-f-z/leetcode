package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2843(t *testing.T) {
	low, high := 1, 100

	ans := countSymmetricIntegers(low, high)

	assert.Check(t, ans == 9)

	low, high = 1200, 1230

	ans = countSymmetricIntegers(low, high)

	assert.Check(t, ans == 4)
}
