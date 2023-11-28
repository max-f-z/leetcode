package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2147(t *testing.T) {
	corridor := "SSPPSPS"

	ans := numberOfWays(corridor)
	assert.Check(t, ans == 3)
}
