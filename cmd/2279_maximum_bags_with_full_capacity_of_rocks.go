package main

import "sort"

//lint:ignore U1000 unused
func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	left := make([]int, len(capacity))

	for i := 0; i < len(capacity); i++ {
		left[i] = capacity[i] - rocks[i]
	}

	sort.Ints(left)

	ans := 0
	for i := 0; i < len(left); i++ {
		if left[i] == 0 {
			ans++
		} else if additionalRocks >= left[i] {
			ans++
			additionalRocks -= left[i]
		}
	}

	return ans
}
