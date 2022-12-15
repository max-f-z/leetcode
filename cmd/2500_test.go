package main

import "testing"

func Test2500(t *testing.T) {
	grid := [][]int{{1, 2, 4}, {3, 3, 1}}

	ans := deleteGreatestValue(grid)
	t.Log(ans)
}
