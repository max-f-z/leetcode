package main

import "strings"

//lint:ignore U1000 unused
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i < len(s)+1; i++ {
		if dp[i-1] == false {
			continue
		}
		for _, v := range wordDict {
			if t := len(v) + i - 1; t < len(s)+1 && strings.HasPrefix(s[i-1:], v) {
				dp[t] = true
			}
		}
	}
	return dp[len(s)]
}
