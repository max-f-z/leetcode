package main

import "math"

func shortestDistance317(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	h, w := len(grid), len(grid[0])
	building := 0
	min := math.MaxInt32
	found := false

	visited := make([][]int, h)

	for i := 0; i < h; i++ {
		visited[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 1 {
				building++
				stack := &[][]int{{i, j, 0}}

				record := make([][]int, h)

				for i := 0; i < h; i++ {
					record[i] = make([]int, w)
				}

				for len(*stack) > 0 {
					x, y, step := (*stack)[0][0], (*stack)[0][1], (*stack)[0][2]
					*stack = (*stack)[1:]

					grid[x][y] -= step

					if x+1 < h && grid[x+1][y] <= 0 && record[x+1][y] == 0 {
						*stack = append(*stack, []int{x + 1, y, step + 1})
						visited[x+1][y]++
						record[x+1][y] = 1
					}

					if x-1 >= 0 && grid[x-1][y] <= 0 && record[x-1][y] == 0 {
						*stack = append(*stack, []int{x - 1, y, step + 1})
						visited[x-1][y]++
						record[x-1][y] = 1
					}

					if y+1 < w && grid[x][y+1] <= 0 && record[x][y+1] == 0 {
						*stack = append(*stack, []int{x, y + 1, step + 1})
						visited[x][y+1]++
						record[x][y+1] = 1
					}

					if y-1 >= 0 && grid[x][y-1] <= 0 && record[x][y-1] == 0 {
						*stack = append(*stack, []int{x, y - 1, step + 1})
						visited[x][y-1]++
						record[x][y-1] = 1
					}
				}
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] < 0 && visited[i][j] == building {
				if min > -grid[i][j] {
					min = -grid[i][j]
					found = true
				}
			}
		}
	}

	if !found {
		return -1
	}

	return min
}
