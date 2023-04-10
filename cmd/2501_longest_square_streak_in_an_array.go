package main

import "sort"

//lint:ignore U1000 unused
func longestSquareStreak(nums []int) int {
	vals := map[int]int{}
	for i := 0; i < len(nums); i++ {
		vals[nums[i]] = 1
	}

	sort.Ints(nums)

	max := -1

	cache := map[int]int{}

	for i := len(nums) - 1; i >= 0; i-- {
		if vals[nums[i]*nums[i]] == 1 {
			cache[nums[i]] = cache[nums[i]*nums[i]] + 1
		} else {
			cache[nums[i]] = 1
		}
		if max < cache[nums[i]] {
			max = cache[nums[i]]
		}
	}

	if max == 1 {
		max = -1
	}

	return max
}
