package main

//lint:ignore U1000 unused
func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	current := 0

	start_index := 0

	for i := 0; i < len(gas); i++ {
		current += gas[i] - cost[i]

		if current < 0 {
			current = 0
			start_index = i + 1
		}

		total += gas[i] - cost[i]
	}

	if total < 0 {
		return -1
	}

	return start_index
}
