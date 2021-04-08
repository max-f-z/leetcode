package main

import (
	"fmt"
	"sort"
)

func nextClosestTime(time string) string {
	origin := []int{}

	res := make([]int, 4)

	for i := 0; i < len(time); i++ {
		if time[i] != ':' {
			tmp := int(time[i] - 48)
			origin = append(origin, tmp)
		}
	}

	sorted := make([]int, 4)

	copy(sorted, origin)
	sort.Ints(sorted)

	var flag bool

	res[3], flag = getNextVal(sorted, 9, origin[3])
	if !flag {
		return fmt.Sprintf("%d%d:%d%d", origin[0], origin[1], origin[2], res[3])
	}

	res[2], flag = getNextVal(sorted, 5, origin[2])
	if !flag {
		return fmt.Sprintf("%d%d:%d%d", origin[0], origin[1], res[2], res[3])
	}

	if origin[0] == 2 {
		res[1], flag = getNextVal(sorted, 4, origin[1])
	} else {
		res[1], flag = getNextVal(sorted, 9, origin[1])
	}

	if !flag {
		return fmt.Sprintf("%d%d:%d%d", origin[0], res[1], res[2], res[3])
	}

	res[0], _ = getNextVal(sorted, 2, origin[0])

	return fmt.Sprintf("%d%d:%d%d", res[0], res[1], res[2], res[3])
}

func getNextVal(digits []int, limit int, target int) (int, bool) {
	for i := 0; i < len(digits); i++ {
		if digits[i] > target && digits[i] <= limit {
			return digits[i], false
		}
	}

	return digits[0], true
}
