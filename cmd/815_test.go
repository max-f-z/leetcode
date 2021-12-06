package main

import "testing"

func Test815(t *testing.T) {
	routes := [][]int{{1, 2, 7}, {3, 6, 7}}
	source := 1
	target := 7

	ans := numBusesToDestination(routes, source, target)

	t.Log(ans)
}
