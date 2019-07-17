package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	if n == 1 && m == 1 {
		if obstacleGrid[0][0] == 1 {
			return 0
		}
		return 1
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = -1
				} else if j > 0 && dp[i][j-1] == -1 {
					dp[i][j] = -1
				} else {
					dp[i][j] = 1
				}
			}

			if j == 0 {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = -1
				} else if i > 0 && dp[i-1][j] == -1 {
					dp[i][j] = -1
				} else {
					dp[i][j] = 1
				}
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = -1
			} else if dp[i-1][j] == -1 && dp[i][j-1] == -1 {
				dp[i][j] = -1
			} else if dp[i-1][j] == -1 || dp[i][j-1] == -1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	// fmt.Println(dp)

	if dp[n-1][m-1] == -1 {
		return 0
	}

	return dp[n-1][m-1]
}
