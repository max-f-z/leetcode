package main

import "sort"

func threeSumSmaller(nums []int, target int) int {
	res := 0

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		l, h := i+1, len(nums)-1
		for l < h {
			if nums[i]+nums[l]+nums[h] < target {
				res += h - l
				l++
			} else {
				h--
			}

		}
	}

	return res
}
