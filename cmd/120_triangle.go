package main

//lint:ignore U1000 unused
func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = triangle[i][j]

			if i == 0 && j == 0 {
				continue
			} else if j == 0 {
				dp[i][j] += dp[i-1][j]
			} else if j == i {
				dp[i][j] += dp[i-1][j-1]
			} else {
				dp[i][j] += min(dp[i-1][j], dp[i-1][j-1])
			}
		}
	}

	m := 10000*n + 1

	for i := 0; i < n; i++ {
		m = min(m, dp[n-1][i])
	}

	return m
}
