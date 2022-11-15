package main

import "testing"

func Test947(t *testing.T) {
	stones := [][]int{{0, 1}, {1, 0}, {1, 1}}

	ans := removeStones(stones)

	t.Log(ans)
}
