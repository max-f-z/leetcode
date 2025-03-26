package main

import (
	"sort"
)

func minOperations2033(grid [][]int, x int) int {
	m := len(grid)
	n := len(grid[0])

	nums := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nums[i*n+j] = grid[i][j]
		}
	}

	sort.Ints(nums)

	median := nums[len(nums)/2]

	operations := 0
	for _, num := range nums {
		if (num-median)%x != 0 {
			return -1
		}
		operations += abs(num-median) / x
	}

	return operations
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
