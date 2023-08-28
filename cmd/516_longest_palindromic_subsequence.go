package main

func longestPalindromeSubseq(s string) int {
	n := len(s)

	srArr := []byte(s)

	for i := 0; i < n/2; i++ {
		srArr[i], srArr[n-1-i] = srArr[n-1-i], srArr[i]
	}
	sr := string(srArr)

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == sr[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j] < dp[i][j-1] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[n][n]
}
