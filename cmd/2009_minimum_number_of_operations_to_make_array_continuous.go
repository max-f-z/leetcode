package main

import "sort"

func minOperations(nums []int) int {
	l := len(nums)
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	dedup := []int{}
	for k := range m {
		dedup = append(dedup, k)
	}

	sort.Ints(dedup)
	max := 0

	for i, v := range dedup {
		num := sort.SearchInts(dedup[i:], v+l)
		if num > max {
			max = num
		}
	}

	return l - max
}
