package main

import "testing"

func Test797(t *testing.T) {
	graph := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}

	ans := allPathsSourceTarget(graph)

	t.Log(ans)
}
