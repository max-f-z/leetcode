package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test650(t *testing.T) {
	ans := minSteps(7)
	assert.Check(t, ans == 7)
}
