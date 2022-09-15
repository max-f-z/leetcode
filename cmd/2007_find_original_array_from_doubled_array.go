package main

import "sort"

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return []int{}
	}

	sort.Ints(changed)

	freq := map[int]int{}

	for _, num := range changed {
		freq[num]++
	}

	ans := []int{}
	for _, num := range changed {
		for freq[num] > 0 {
			if freq[num*2] > 0 {
				freq[num]--
				freq[num*2]--
				ans = append(ans, num)
			} else {
				return []int{}
			}
		}
	}

	return ans
}
