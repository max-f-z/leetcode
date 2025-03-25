package main

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	horizontal := [][]int{}
	vertical := [][]int{}

	for _, r := range rectangles {
		horizontal = append(horizontal, []int{r[0], r[2]})
		vertical = append(vertical, []int{r[1], r[3]})
	}

	sort.Slice(horizontal, func(i, j int) bool {
		return horizontal[i][0] < horizontal[j][0]
	})

	sort.Slice(vertical, func(i, j int) bool {
		return vertical[i][0] < vertical[j][0]
	})

	bar := horizontal[0][1]
	cnt := 0

	for i := 1; i < len(horizontal); i++ {
		if bar <= horizontal[i][0] {
			cnt++
			bar = horizontal[i][1]
		} else {
			bar = max(bar, horizontal[i][1])
		}
	}
	if cnt >= 2 {
		return true
	}

	bar = vertical[0][1]
	cnt = 0

	for i := 1; i < len(vertical); i++ {
		if bar <= vertical[i][0] {
			cnt++
			bar = vertical[i][1]
		} else {
			bar = max(bar, vertical[i][1])
		}
	}

	return cnt >= 2
}
