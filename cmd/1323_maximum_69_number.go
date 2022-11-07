package main

import (
	"strconv"
)

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	flip := false
	ans := 0
	for _, v := range str {
		if v == '6' && !flip {
			flip = true
			ans = ans*10 + 9
		} else {
			tmp := 0
			if v == '6' {
				tmp = 6
			} else {
				tmp = 9
			}
			ans = ans*10 + tmp
		}
	}

	return ans
}
