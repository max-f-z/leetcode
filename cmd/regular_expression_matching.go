package main

/*
https://www.youtube.com/watch?v=KN22ZEpRTFY
*/
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for j := 1; j < len(p)+1; j++ {
		if string(p[j-1]) == "*" {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if string(s[i-1]) == string(p[j-1]) || string(p[j-1]) == "." {
				dp[i][j] = dp[i-1][j-1]
			}

			if j > 1 && string(p[j-1]) == "*" {
				if string(p[j-2]) != string(s[i-1]) && string(p[j-2]) != "." {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-1] || dp[i][j-2] || dp[i-1][j]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}
