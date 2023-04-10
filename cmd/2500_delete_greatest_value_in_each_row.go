package main

import "sort"

//lint:ignore U1000 unused
func deleteGreatestValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		max := -1
		for j := 0; j < m; j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		ans += max
	}

	return ans
}
