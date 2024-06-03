package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2486(t *testing.T) {
	a := "coaching"
	b := "coding"

	ans := appendCharacters(a, b)

	assert.Check(t, ans == 4)
}
