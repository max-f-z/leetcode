package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test135(t *testing.T) {
	ratings := []int{1, 0, 2}

	ans := candy(ratings)

	assert.Check(t, ans == 5)
}
