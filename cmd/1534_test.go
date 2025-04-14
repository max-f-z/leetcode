package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1534(t *testing.T) {
	arr := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3
	ans := countGoodTriplets(arr, a, b, c)
	assert.Check(t, ans == 4)

	arr = []int{7, 3, 7, 3, 12, 1, 12, 2, 3}
	a, b, c = 5, 8, 1
	ans = countGoodTriplets(arr, a, b, c)
	assert.Check(t, ans == 12)
}
