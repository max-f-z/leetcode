package main

import (
	"strconv"
)

func findMissingRanges(arr []int, lower int, upper int) []string {
	if len(arr) == 0 {
		return []string{range2str(lower, upper)}
	}

	current := lower
	out := make([]string, 0)

	for i := 0; i < len(arr); i++ {
		switch {
		case arr[i] < current:
			// ignore

		case arr[i] == current:
			// same
			current++

		case i > 0 && arr[i] == current:
			// merge intervals
			current = arr[i] + 1

		case arr[i] > current && arr[i] <= upper:
			// in interval
			start := current
			stop := arr[i] - 1

			out = append(out, range2str(start, stop))
			current = arr[i] + 1
		}
	}

	if current <= upper {
		out = append(out, range2str(current, upper))
	}

	return out
}

func range2str(a, b int) string {
	if a > b {
		return strconv.Itoa(b)
	}

	if a == b {
		return strconv.Itoa(a)
	}
	return strconv.Itoa(a) + "->" + strconv.Itoa(b)
}
