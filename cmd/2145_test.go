package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2145(t *testing.T) {
	differences := []int{1, -3, 4}
	lower, upper := 1, 6
	ans := numberOfArrays(differences, lower, upper)
	assert.Check(t, ans == 2)

	differences = []int{-40}
	lower, upper = -46, 53
	ans = numberOfArrays(differences, lower, upper)
	assert.Check(t, ans == 60)
}
