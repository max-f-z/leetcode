package main

import "math"

func shortestDistance(words []string, word1 string, word2 string) int {
	if len(words) == 2 {
		return 1
	}

	min := math.MaxInt64
	idx1 := -1
	idx2 := -1
	for i, v := range words {
		if v == word1 {
			idx1 = i
		}
		if v == word2 {
			idx2 = i
		}

		if idx1 != -1 && idx2 != -1 && int(math.Abs(float64(idx1-idx2))) < min {
			min = int(math.Abs(float64(idx1 - idx2)))
		}
	}

	return min
}
