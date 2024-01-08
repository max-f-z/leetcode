package main

//lint:ignore U1000 unused
func minDistance(word1 string, word2 string) int {
	n, m := len(word1)+1, len(word2)+1

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = i
		for j := 0; j < m; j++ {
			dp[0][j] = j
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
			}
		}
	}

	return dp[n-1][m-1]
}

// func min(a ...int) int {
// 	min := math.MaxInt64
// 	for _, v := range a {
// 		if v < min {
// 			min = v
// 		}
// 	}
// 	return min
// }
