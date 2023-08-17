package main

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
				continue
			}

			min := m*n + 1
			if i-1 >= 0 && min > dp[i-1][j]+1 {
				min = dp[i-1][j] + 1
			}

			if j-1 >= 0 && min > dp[i][j-1]+1 {
				min = dp[i][j-1] + 1
			}

			dp[i][j] = min
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			min := dp[i][j]

			if i+1 < m && min > dp[i+1][j]+1 {
				min = dp[i+1][j] + 1
			}

			if j+1 < n && min > dp[i][j+1]+1 {
				min = dp[i][j+1] + 1
			}

			dp[i][j] = min
		}
	}

	return dp
}
