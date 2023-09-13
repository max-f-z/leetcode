package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1647(t *testing.T) {
	ans := minDeletions("bbcebab")

	assert.Check(t, ans == 2)
}
