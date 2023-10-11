package main

import "testing"

func Test2251(t *testing.T) {
	flowers := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	people := []int{2, 3, 7, 11}

	ans := fullBloomFlowers(flowers, people)

	t.Log(ans)
}
