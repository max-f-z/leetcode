package main

import (
	"math"
)

//lint:ignore U1000 unused
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		if coin < amount+1 {
			dp[coin] = 1
		}
	}

	for i := 2; i < amount+1; i++ {
		for _, coin := range coins {
			if i > coin && dp[i] > (dp[i-coin]+1) {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
