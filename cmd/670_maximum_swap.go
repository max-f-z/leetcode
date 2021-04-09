package main

import "math"

func maximumSwap(num int) int {
	reverse := intToSlice(num)

	memo := make([][]int, len(reverse))

	max := -1
	idx := -1

	for i := 0; i < len(reverse); i++ {
		if max < reverse[i] {
			max = reverse[i]
			idx = i
			memo[len(reverse)-1-i] = []int{reverse[i], idx}
		} else {
			memo[len(reverse)-1-i] = []int{max, idx}
		}
	}

	res := 0

	for i := 0; i < len(memo); i++ {
		if reverse[len(reverse)-1-i] < memo[i][0] {
			reverse[len(reverse)-1-i], reverse[memo[i][1]] = reverse[memo[i][1]], reverse[len(reverse)-1-i]
			break
		}
	}

	for i := 0; i < len(reverse); i++ {
		res = res + reverse[i]*int(math.Pow10(i))
	}

	return res
}

func intToSlice(num int) []int {
	res := []int{}

	for num > 0 {
		tmp := num % 10
		res = append(res, tmp)
		num = (num - tmp) / 10
	}

	return res
}
