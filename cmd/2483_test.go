package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2483(t *testing.T) {
	customers := "YYYY"

	ans := bestClosingTime(customers)

	assert.Check(t, ans == 4)
}
