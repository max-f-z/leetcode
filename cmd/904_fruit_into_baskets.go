package main

func totalFruit(fruits []int) int {
	taken := map[int]int{}

	p1, p2 := 0, 0

	max := 0

	for p2 < len(fruits) {
		taken[fruits[p2]]++

		for len(taken) > 2 {
			taken[fruits[p1]]--
			if taken[fruits[p1]] == 0 {
				delete(taken, fruits[p1])
			}
			p1++
		}

		if p2-p1+1 > max {
			max = p2 - p1 + 1
		}

		p2++
	}

	return max
}
