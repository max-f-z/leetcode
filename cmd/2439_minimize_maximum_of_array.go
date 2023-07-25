package main

func minimizeArrayValue(nums []int) int {
	sum := 0
	prefixSum := make([]int, len(nums))

	for idx, v := range nums {
		sum += v
		prefixSum[idx] = sum
	}

	dp := make([]int, len(nums))

	dp[0] = prefixSum[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1]

		if prefixSum[i] > dp[i]*(i+1) {
			if prefixSum[i]%(i+1) == 0 {
				dp[i] = prefixSum[i] / (i + 1)
			} else {
				dp[i] = prefixSum[i]/(i+1) + 1
			}
		}
	}

	return dp[len(nums)-1]
}
