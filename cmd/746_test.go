package main

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	costs := []int{10, 15, 20}

	ans := minCostClimbingStairs(costs)

	t.Log(ans)
}
