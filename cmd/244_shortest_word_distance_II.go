package main

import "math"

type WordDistance struct {
	data []string
}

func Constructor2(words []string) WordDistance {
	wd := WordDistance{data: words}
	return wd
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	if len(this.data) == 2 {
		return 1
	}

	min := math.MaxInt64
	idx1 := -1
	idx2 := -1
	for i, v := range this.data {
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
