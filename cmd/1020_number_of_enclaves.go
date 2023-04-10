package main

//lint:ignore U1000 unused
func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && grid[i][j] == 1 && !visited[i][j] {
				checkIsland(&grid, &visited, i, j, m, n)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count++
			}
		}
	}

	return count
}

func checkIsland(grid *[][]int, visited *[][]bool, i, j, m, n int) {
	if i < 0 || i == m || j < 0 || j == n || (*grid)[i][j] != 1 || (*visited)[i][j] {
		return
	}

	(*visited)[i][j] = true

	checkIsland(grid, visited, i-1, j, m, n)
	checkIsland(grid, visited, i+1, j, m, n)
	checkIsland(grid, visited, i, j-1, m, n)
	checkIsland(grid, visited, i, j+1, m, n)
}
