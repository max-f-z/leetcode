package main

import "testing"

func Test1129(t *testing.T) {
	n := 5
	red := [][]int{{2, 2}, {0, 1}, {0, 3}, {0, 0}, {0, 4}, {2, 1}, {2, 0}, {1, 4}, {3, 4}}
	blue := [][]int{{1, 3}, {0, 0}, {0, 3}, {4, 2}, {1, 0}}

	ans := shortestAlternatingPaths(n, red, blue)

	t.Log(ans)
}
