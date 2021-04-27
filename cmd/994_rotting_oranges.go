package main

func orangesRotting(grid [][]int) int {
	queue := [][]int{}

	m, n := len(grid), len(grid[0])

	cnt := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}

			if grid[i][j] == 1 {
				cnt++
			}
		}
	}

	step := 0
	for len(queue) > 0 {
		newQ := [][]int{}

		for _, v := range queue {
			if v[0]+1 >= 0 && v[0]+1 < m && grid[v[0]+1][v[1]] == 1 {
				newQ = append(newQ, []int{v[0] + 1, v[1]})
				grid[v[0]+1][v[1]] = 2
				cnt--
			}

			if v[0]-1 >= 0 && v[0]-1 < m && grid[v[0]-1][v[1]] == 1 {
				newQ = append(newQ, []int{v[0] - 1, v[1]})
				grid[v[0]-1][v[1]] = 2
				cnt--
			}

			if v[1]+1 >= 0 && v[1]+1 < n && grid[v[0]][v[1]+1] == 1 {
				newQ = append(newQ, []int{v[0], v[1] + 1})
				grid[v[0]][v[1]+1] = 2
				cnt--
			}

			if v[1]-1 >= 0 && v[1]-1 < n && grid[v[0]][v[1]-1] == 1 {
				newQ = append(newQ, []int{v[0], v[1] - 1})
				grid[v[0]][v[1]-1] = 2
				cnt--
			}
		}

		queue = newQ
		if len(queue) != 0 {
			step++
		}
	}

	if cnt != 0 {
		return -1
	}

	return step
}
