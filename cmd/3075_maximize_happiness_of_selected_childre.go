package main

import "slices"

func maximumHappinessSum(happiness []int, k int) int64 {
	slices.Sort(happiness)

	n := len(happiness)
	var maximum int64
	var offset int64
	for i := 0; i < k; i++ {
		val := int64(happiness[n-1-i]) - offset

		if val < 0 {
			val = 0
		}

		maximum += val

		offset += 1
	}

	return maximum
}
