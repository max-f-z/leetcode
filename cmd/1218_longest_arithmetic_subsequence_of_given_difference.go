package main

func longestSubsequence(arr []int, difference int) int {
	dp := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		dp[i] = 1
	}
	max := 0

	cnt := map[int]int{}

	for i := 0; i < len(arr); i++ {
		if value, ok := cnt[arr[i]-difference]; ok {
			cnt[arr[i]] = value + 1
		} else {
			cnt[arr[i]] = 1
		}

		if cnt[arr[i]] > max {
			max = cnt[arr[i]]
		}
	}

	return max
}
