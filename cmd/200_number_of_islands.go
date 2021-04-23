package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var val byte
	val = '1'

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				val++
				numIslandsHelper(grid, val, i, j, m, n)
			}
		}
	}

	return int(val - 1 - 48)
}

func numIslandsHelper(grid [][]byte, val byte, x int, y int, m int, n int) {
	if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
		grid[x][y] = val
	} else {
		return
	}

	numIslandsHelper(grid, val, x+1, y, m, n)
	numIslandsHelper(grid, val, x-1, y, m, n)
	numIslandsHelper(grid, val, x, y+1, m, n)
	numIslandsHelper(grid, val, x, y-1, m, n)
}
