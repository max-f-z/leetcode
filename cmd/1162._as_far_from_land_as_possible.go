package main

func maxDistance(grid [][]int) int {
	step := 0

	n := len(grid)

	queue := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	if len(queue) == 0 || len(queue) == n*n {
		return -1
	}

	for len(queue) > 0 {
		newQ := [][]int{}
		for _, pair := range queue {
			x, y := pair[0], pair[1]

			if x+1 >= 0 && x+1 < n && grid[x+1][y] == 0 {
				newQ = append(newQ, []int{x + 1, y})
				grid[x+1][y] = 1
			}

			if x-1 >= 0 && x-1 < n && grid[x-1][y] == 0 {
				newQ = append(newQ, []int{x - 1, y})
				grid[x-1][y] = 1
			}

			if y+1 >= 0 && y+1 < n && grid[x][y+1] == 0 {
				newQ = append(newQ, []int{x, y + 1})
				grid[x][y+1] = 1
			}

			if y-1 >= 0 && y-1 < n && grid[x][y-1] == 0 {
				newQ = append(newQ, []int{x, y - 1})
				grid[x][y-1] = 1
			}
		}

		queue = newQ

		if len(newQ) > 0 {
			step++
		}
	}

	return step
}
