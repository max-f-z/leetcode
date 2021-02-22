package main

func findLength(A []int, B []int) int {
	a, b := len(A), len(B)

	dp := make([][]int, a)
	for i := 0; i < a; i++ {
		dp[i] = make([]int, b)
	}

	res := 0

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if A[i] == B[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1

					if res < dp[i][j] {
						res = dp[i][j]
					}
				}
			}
		}
	}

	return res

}
