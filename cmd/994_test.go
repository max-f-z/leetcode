package main

import "testing"

func Test994(t *testing.T) {
	m := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}

	n := orangesRotting(m)

	t.Log(n)
}
