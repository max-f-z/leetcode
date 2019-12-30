package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		numOnes := 0
		for _, c := range str {
			if string(c) == "1" {
				numOnes++
			}
		}
		numZeros := len(str) - numOnes

		for i := m; i >= 0; i-- {
			for j := n; j >= 0; j-- {
				if i >= numZeros && j >= numOnes {
					temp := dp[i-numZeros][j-numOnes] + 1
					if temp > dp[i][j] {
						dp[i][j] = temp
					}
				}
			}
		}
	}
	return dp[m][n]
}
