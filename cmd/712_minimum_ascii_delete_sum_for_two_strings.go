package main

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}

	for i := 1; i < n+1; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if dp[i-1][j]+int(s1[i-1]) < dp[i][j-1]+int(s2[j-1]) {
					dp[i][j] = dp[i-1][j] + int(s1[i-1])
				} else {
					dp[i][j] = dp[i][j-1] + int(s2[j-1])
				}
			}
		}
	}

	return dp[m][n]
}
