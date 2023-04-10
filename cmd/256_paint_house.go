package main

//lint:ignore U1000 unused
func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	dp := [][]int{}

	n := len(costs)

	dp = append(dp, []int{0, 0, 0})

	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]

	for i := 1; i < len(costs); i++ {
		dp = append(dp, []int{0, 0, 0})
		dp[i][0] = minCostMin(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = minCostMin(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = minCostMin(dp[i-1][1], dp[i-1][0]) + costs[i][2]
	}

	return minCostMin3(dp[n-1][0], dp[n-1][1], dp[n-1][2])
}

func minCostMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostMin3(a, b, c int) int {
	return minCostMin(minCostMin(a, b), c)
}
