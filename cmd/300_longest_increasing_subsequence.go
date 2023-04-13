package main

//lint:ignore U1000 unused
func lengthOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	dp[0] = 1

	ans := 1

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		dp[i] = max(dp[i], 1)

		ans = max(dp[i], ans)
	}

	return ans
}
