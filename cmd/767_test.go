package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test767(t *testing.T) {
	m := reorganizeString("aab")

	assert.Check(t, m == "aba")
}
