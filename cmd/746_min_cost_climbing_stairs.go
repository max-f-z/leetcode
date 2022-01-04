package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < n+1; i++ {
		oneStep := dp[i-1] + cost[i-1]
		twoStep := dp[i-2] + cost[i-2]

		if oneStep < twoStep {
			dp[i] = oneStep
		} else {
			dp[i] = twoStep
		}
	}

	return dp[n]
}
