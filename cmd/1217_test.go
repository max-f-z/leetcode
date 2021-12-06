package main

import "testing"

func Test1217(t *testing.T) {
	position := []int{1, 2, 3}

	cost := minCostToMoveChips(position)

	t.Log(cost)
}
