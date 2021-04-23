package main

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	max := 0

	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - 48)
		if dp[i][0] == 1 && max == 0 {
			max = 1
		}
	}

	for i := 0; i < n; i++ {
		dp[0][i] = int(matrix[0][i] - 48)
		if dp[0][i] == 1 && max == 0 {
			max = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				val := minVal(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
				if val == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = val + 1
				}
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max
}

func minVal(a ...int) int {
	val := a[0]

	for i := 0; i < len(a); i++ {
		if val > a[i] {
			val = a[i]
		}
	}

	return val
}
