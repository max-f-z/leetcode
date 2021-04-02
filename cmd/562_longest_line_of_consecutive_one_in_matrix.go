package main

func longestLine(M [][]int) int {
	max := make([]int, 4)

	dp := make([][][]int, 4)

	l := len(M)

	if l == 0 {
		return 0
	}

	w := len(M[0])

	// init
	for i := 0; i < 4; i++ {
		dp[i] = make([][]int, l)
		for j := 0; j < l; j++ {
			dp[i][j] = make([]int, w)
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			// left - right

			if M[i][j] == 1 {
				if j-1 >= 0 {
					dp[0][i][j] = dp[0][i][j-1] + 1
				} else {
					dp[0][i][j] = 1
				}
			} else {
				dp[0][i][j] = 0
			}

			if dp[0][i][j] > max[0] {
				max[0] = dp[0][i][j]
			}

			// top - down

			if M[i][j] == 1 {
				if i-1 >= 0 {
					dp[1][i][j] = dp[1][i-1][j] + 1
				} else {
					dp[1][i][j] = 1
				}
			} else {
				dp[1][i][j] = 0
			}

			if dp[1][i][j] > max[1] {
				max[1] = dp[1][i][j]
			}

			// top left - right down

			if M[i][j] == 1 {
				if i-1 >= 0 && j-1 >= 0 {
					dp[2][i][j] = dp[2][i-1][j-1] + 1
				} else {
					dp[2][i][j] = 1
				}
			} else {
				dp[2][i][j] = 0
			}

			if dp[2][i][j] > max[2] {
				max[2] = dp[2][i][j]
			}

			// top right - left down

			if M[l-1-i][j] == 1 {
				if l-1-i+1 < l && j-1 >= 0 {
					dp[3][l-1-i][j] = dp[3][l-1-i+1][j-1] + 1
				} else {
					dp[3][l-1-i][j] = 1
				}
			} else {
				dp[3][l-1-i][j] = 0
			}

			if dp[3][l-1-i][j] > max[3] {
				max[3] = dp[3][l-1-i][j]
			}
		}
	}

	res := 0

	for i := 0; i < 4; i++ {
		if res < max[i] {
			res = max[i]
		}
	}

	return res
}
