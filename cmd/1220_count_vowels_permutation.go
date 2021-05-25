package main

func countVowelPermutation(n int) int {
	dp := []int{1, 1, 1, 1, 1}

	mod := 1000000007

	if n == 1 {
		return 5
	}

	for i := 1; i < n; i++ {
		tmp := make([]int, 5)

		tmp[0] = (dp[1] + dp[2] + dp[4]) % mod
		tmp[1] = (dp[0] + dp[2]) % mod
		tmp[2] = (dp[1] + dp[3]) % mod
		tmp[3] = dp[2] % mod
		tmp[4] = (dp[2] + dp[3]) % mod

		dp = tmp
	}

	return (dp[0] + dp[1] + dp[2] + dp[3] + dp[4]) % mod
}
