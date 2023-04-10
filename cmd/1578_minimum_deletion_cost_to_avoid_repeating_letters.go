package main

//lint:ignore U1000 unused
func minCost1578(s string, cost []int) int {
	val := 0

	var char byte
	char = '*'
	idx := -1

	for i := 0; i < len(s); i++ {
		if char != s[i] {
			char = s[i]
			idx = i
		} else {
			if cost[i] > cost[idx] {
				val += cost[idx]
				idx = i
			} else {
				val += cost[i]
			}
		}
	}

	return val
}
