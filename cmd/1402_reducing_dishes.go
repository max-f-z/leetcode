package main

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)

	n := len(satisfaction)

	sum := 0
	previousScore := 0
	for i := n - 1; i >= 0; i-- {
		sum += satisfaction[i]
		if sum > 0 {
			previousScore = previousScore + sum
		} else {
			break
		}
	}

	return previousScore
}
