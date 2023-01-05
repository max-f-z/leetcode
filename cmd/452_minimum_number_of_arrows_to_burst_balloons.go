package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i int, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	ans := 0

	start, end := points[0][0], points[0][1]
	for idx, v := range points {
		if idx == 0 {
			continue
		}

		if v[0] > end || v[1] < start {
			ans++
			start = v[0]
			end = v[1]
		}

		if v[0] > start {
			start = v[0]
		}

		if v[1] < end {
			end = v[1]
		}

		continue
	}

	ans++

	return ans
}
