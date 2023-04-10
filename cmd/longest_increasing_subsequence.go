package main

//lint:ignore U1000 unused
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for j := range nums {
		max := 0
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
		if j == 0 {
			dp[j] = 1
		} else {
			dp[j] = max + 1
		}
	}

	// fmt.Println(dp)

	max := 0
	for _, m := range dp {
		if m > max {
			max = m
		}
	}

	return max
}
