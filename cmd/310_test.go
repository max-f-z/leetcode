package main

import (
	"testing"
)

func Test310(t *testing.T) {

	edges := [][]int{{1, 0}, {2, 1}, {3, 1}, {2, 4}, {5, 3}, {4, 6}}

	ans := findMinHeightTrees(7, edges)

	t.Log(ans)
}
