package main

//lint:ignore U1000 unused
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	if len(nums) == 1 {
		return nums[0]
	}
	dp[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}
