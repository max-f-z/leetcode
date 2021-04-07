package main

func countSubstrings(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 1
		}
	}

	for l := 3; l <= len(s); l++ {
		for i := 0; i < len(s)-l+1; i++ {
			if s[i] == s[i+l-1] && dp[i+1][i+l-2] == 1 {
				dp[i][i+l-1] = 1
			}
		}
	}

	cnt := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if dp[i][j] == 1 {
				cnt++
			}
		}
	}

	return cnt
}
