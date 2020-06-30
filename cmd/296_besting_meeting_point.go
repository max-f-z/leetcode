package main

import "sort"

func minTotalDistance(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	x := []int{}
	y := []int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 1 {
				x = append(x, i)
				y = append(y, j)
			}
		}
	}

	sort.Ints(x)
	sort.Ints(y)

	sumX := 0
	sumY := 0

	l := 0
	r := len(x) - 1
	for l < r {
		sumX += x[r] - x[l]
		l++
		r--
	}
	l = 0
	r = len(y) - 1
	for l < r {
		sumY += y[r] - y[l]
		l++
		r--
	}

	return sumX + sumY
}
