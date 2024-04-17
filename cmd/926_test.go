package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test926(t *testing.T) {
	ans := minFlipsMonoIncr("00011000")

	assert.Check(t, ans == 2)
}
