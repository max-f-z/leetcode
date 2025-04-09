package main

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i := 0; i < len(citations); i++ {
		if citations[i] >= (i + 1) {
			continue
		} else {
			return i
		}
	}
	return len(citations)
}
