package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test168(t *testing.T) {
	ans := convertToTitle(26)

	assert.Check(t, ans == "Z")

	ans = convertToTitle(1)

	assert.Check(t, ans == "A")

	ans = convertToTitle(28)

	assert.Check(t, ans == "AB")

	ans = convertToTitle(701)

	assert.Check(t, ans == "ZY")

	ans = convertToTitle(52)

	assert.Check(t, ans == "AZ")
}
