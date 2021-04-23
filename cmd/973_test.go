package main

import "testing"

func Test973(t *testing.T) {
	points := [][]int{{1, 3}, {-2, 2}, {2, -2}}
	k := 2

	val := kClosest(points, k)

	t.Log(val)
}
