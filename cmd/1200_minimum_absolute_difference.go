package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	ans := [][]int{}

	sort.Ints(arr)
	min := math.MaxInt32

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			ans = [][]int{{arr[i-1], arr[i]}}
			min = arr[i] - arr[i-1]
		} else if arr[i]-arr[i-1] == min {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}
