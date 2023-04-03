package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	p1, p2 := 0, len(people)-1

	ans := 0

	for p1 <= p2 {
		if people[p1]+people[p2] <= limit {
			p1++
			p2--
			ans++
		} else {
			p2--
			ans++
		}
	}

	return ans
}
