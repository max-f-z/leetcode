package main

import (
	"math"
)

func myAtoi(str string) int {
	var ans int64 = 0
	var sign int64 = 1
	start := false

	for _, b := range str {
		if b <= '9' && b >= '0' {
			if !start {
				start = true
			}
			ans = ans*10 + int64(b) - int64('0')
			if ans*sign > math.MaxInt32 {
				return math.MaxInt32
			} else if ans*sign < math.MinInt32 {
				return math.MinInt32
			}
		} else if !start && b == '+' {
			start = true
		} else if !start && b == '-' {
			start = true
			sign = -1
		} else if !start && b == ' ' {
			continue
		} else {
			break
		}
	}
	return int(ans * sign)
}
