package main

import "fmt"

//lint:ignore U1000 unused
func uniquePaths(m int, n int) int {

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		if i == 0 {
			for v := range dp[0] {
				dp[0][v] = 1
			}
		}
		dp[i][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	fmt.Println(dp)

	return dp[n-1][m-1]
}
