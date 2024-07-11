package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1642(t *testing.T) {
	heights := []int{4, 2, 7, 6, 9, 14, 12}
	ans := furthestBuilding(heights, 5, 1)
	assert.Check(t, ans == 4)

	heights = []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	ans = furthestBuilding(heights, 10, 2)
	assert.Check(t, ans == 7)

	heights = []int{14, 3, 19, 3}
	ans = furthestBuilding(heights, 17, 0)
	assert.Check(t, ans == 3)
}
