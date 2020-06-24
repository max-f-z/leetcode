package main

import "math"

// dp 公式同 256 paint house
// 对于每一行，保留最小的两个cost

func minCostII(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}

	n := len(costs)
	k := len(costs[0])

	dp := [][]int{}

	dp = append(dp, make([]int, k))

	min := make([][]int, n)
	for i := 0; i < n; i++ {
		min[i] = []int{math.MaxInt32, math.MaxInt32}
	}

	for i := 0; i < k; i++ {
		dp[0][i] = costs[0][i]
		if dp[0][i] < min[0][0] {
			min[0][0], min[0][1] = dp[0][i], min[0][0]
			continue
		}

		if dp[0][i] < min[0][1] {
			min[0][1] = dp[0][i]
		}
	}

	for i := 1; i < n; i++ {
		dp = append(dp, make([]int, k))
		for j := 0; j < k; j++ {
			if dp[i-1][j] != min[i-1][0] {
				dp[i][j] = costs[i][j] + min[i-1][0]
			} else {
				dp[i][j] = costs[i][j] + min[i-1][1]
			}

			if dp[i][j] < min[i][0] {
				min[i][0], min[i][1] = dp[i][j], min[i][0]
				continue
			}

			if dp[i][j] < min[i][1] {
				min[i][1] = dp[i][j]
			}
		}
	}

	return min[n-1][0]
}
