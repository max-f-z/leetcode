package main

import "math"

//lint:ignore U1000 unused
func editDistance(word1, word2 string) int {
	m, n := len(word1)+1, len(word2)+1

	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	/*
		dp[i][0] = i
		dp[0][j] = j

		dp[m-1][n-1] should be the expected value
	*/
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = i
	}

	for j := 0; j < n; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min := math.MaxInt
				if min > dp[i-1][j] {
					min = dp[i-1][j]
				}
				if min > dp[i][j-1] {
					min = dp[i][j-1]
				}
				if min > dp[i-1][j-1] {
					min = dp[i-1][j-1]
				}
				dp[i][j] = min + 1
			}
		}
	}

	return dp[m-1][n-1]
}
