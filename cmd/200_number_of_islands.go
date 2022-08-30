package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				numIslandsHelper(&grid, i, j, m, n)
			}
		}
	}

	return res
}

func numIslandsHelper(grid *[][]byte, x int, y int, m int, n int) {
	if x >= 0 && x < m && y >= 0 && y < n && (*grid)[x][y] == '1' {
		(*grid)[x][y] = '0'
	} else {
		return
	}
	// (*grid)[x][y] = '0'
	numIslandsHelper(grid, x+1, y, m, n)
	numIslandsHelper(grid, x-1, y, m, n)
	numIslandsHelper(grid, x, y+1, m, n)
	numIslandsHelper(grid, x, y-1, m, n)
}
