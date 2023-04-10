package main

import (
	"fmt"
	"strconv"
)

//lint:ignore U1000 unused
func numDecodings(s string) int {
	if s == "0" || string(s[0]) == "0" {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < len(s)+1; i++ {
		if string(s[i-1]) != "0" {
			dp[i] = dp[i-1]
			if string(s[i-2]) != "0" {
				val, _ := strconv.Atoi(string(s[i-2 : i]))
				if val <= 26 {
					dp[i] += dp[i-2]
				}
			}
		} else {
			if string(s[i-2]) != "1" && string(s[i-2]) != "2" {
				return 0
			}
			dp[i] = dp[i-2]
		}
	}

	fmt.Println(dp)

	return dp[len(dp)-1]
}
