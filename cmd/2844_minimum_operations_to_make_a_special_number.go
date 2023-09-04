package main

import "strings"

func minimumOperations(num string) int {
	zeroIdx := strings.LastIndex(num, "0")
	fiveIdx := strings.LastIndex(num, "5")

	ans := len(num)

	if fiveIdx != -1 {
		for i := fiveIdx - 1; i >= 0; i-- {
			if num[i] == '2' || num[i] == '7' {
				val := len(num) - 1 - fiveIdx
				val += fiveIdx - i - 1
				if val < ans {
					ans = val
				}
			}
		}
	}

	if zeroIdx != -1 {
		for i := zeroIdx - 1; i >= 0; i-- {
			if num[i] == '5' || num[i] == '0' {
				val := len(num) - 1 - zeroIdx
				val += zeroIdx - i - 1
				if val < ans {
					ans = val
				}
			}
		}
		if ans > len(num)-1 {
			ans = len(num) - 1
		}
	}

	return ans
}
