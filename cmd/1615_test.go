package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1615(t *testing.T) {
	roads := [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}

	ans := maximalNetworkRank(4, roads)

	assert.Check(t, ans == 4)

	roads = [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}

	ans = maximalNetworkRank(5, roads)

	assert.Check(t, ans == 5)

	roads = [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}

	ans = maximalNetworkRank(8, roads)

	assert.Check(t, ans == 5)
}
