package main

//lint:ignore U1000 unused
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = minInt(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[n-1][m-1]
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
