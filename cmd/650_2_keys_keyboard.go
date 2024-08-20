package main

func minSteps(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i < n+1; i++ {
		dp[i] = i
	}

	for i := 2; i < n+1; i++ {
		for j := 1; j < i/2; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
			}
		}
	}

	return dp[n]
}
