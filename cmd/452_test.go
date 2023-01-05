package main

import "testing"

func Test452(t *testing.T) {
	vals := [][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}

	ans := findMinArrowShots(vals)

	t.Log(ans)
}
