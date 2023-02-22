package main

func shipWithinDays(weights []int, days int) int {
	max := 0

	for i := 0; i < len(weights); i++ {
		if max < weights[i] {
			max = weights[i]
		}
	}

	l := max
	r := max * len(weights)

	for l < r {
		mid := (l + r) / 2

		if canShip(weights, mid, days) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func canShip(weights []int, limit int, daysLimit int) bool {
	sum := 0
	days := 1
	for i := 0; i < len(weights); i++ {
		if sum+weights[i] <= limit {
			sum += weights[i]
		} else {
			days++
			sum = weights[i]
		}
	}

	return days <= daysLimit
}
