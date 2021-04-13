package main

import "testing"

func TestDoubleLRU(t *testing.T) {
	ops := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}

	res := LRU(ops, 3)

	t.Log(res)
}
