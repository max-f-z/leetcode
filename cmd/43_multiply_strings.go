package main

import (
	"strconv"
)

func multiply43(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	nums1 := make([]int, len(num1))
	nums2 := make([]int, len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		nums1[len(num1)-1-i], _ = strconv.Atoi(string(num1[i]))
	}

	for i := len(num2) - 1; i >= 0; i-- {
		nums2[len(num2)-1-i], _ = strconv.Atoi(string(num2[i]))
	}

	res := make([]int, len(num1)+len(num2))

	for i, v := range nums2 {
		for j, w := range nums1 {
			res[i+j] += v * w
		}
	}

	tmp := 0
	for i := range res {
		res[i] += tmp
		if res[i] >= 10 {
			tmp = res[i] / 10
			res[i] = res[i] - tmp*10
		} else {
			tmp = 0
		}
	}

	resStr := ""
	start := false
	for i := len(res) - 1; i >= 0; i-- {
		if !start && res[i] != 0 {
			resStr = resStr + strconv.Itoa(res[i])
			start = true
			continue
		}

		if start {
			resStr = resStr + strconv.Itoa(res[i])
		}
	}

	return resStr
}
