package main

import "testing"

func Test1443(t *testing.T) {
	edges := [][]int{{0, 2}, {0, 3}, {1, 2}}

	n := 7

	hasApples := []bool{false, true, false, false}

	ans := minTime(n, edges, hasApples)

	t.Log(ans)
}
