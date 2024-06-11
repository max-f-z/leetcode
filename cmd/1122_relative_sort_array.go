package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	idx := map[int]int{}

	for i := 0; i < len(arr2); i++ {
		idx[arr2[i]] = i + 1
	}

	sort.Slice(arr1, func(i, j int) bool {
		if idx[arr1[i]] != 0 && idx[arr1[j]] != 0 {
			return idx[arr1[i]] < idx[arr1[j]]
		}

		if idx[arr1[i]] == 0 && idx[arr1[j]] != 0 {
			return false
		}

		if idx[arr1[j]] == 0 && idx[arr1[i]] != 0 {
			return true
		}

		return arr1[i] < arr1[j]
	})

	return arr1
}
