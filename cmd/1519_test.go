package main

import "testing"

func Test1519(t *testing.T) {
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	n := 7
	labels := "abaedcd"

	ans := countSubTrees(n, edges, labels)

	t.Log(ans)
}
