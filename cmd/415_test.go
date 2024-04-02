package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test415(t *testing.T) {
	ans := addStrings("1", "9")

	assert.Check(t, ans == "10")
}
