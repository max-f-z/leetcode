package main

//lint:ignore U1000 unused
func minimumTime(time []int, totalTrips int) int64 {
	minTime := 10000001
	for i := 0; i < len(time); i++ {
		if minTime > time[i] {
			minTime = time[i]
		}
	}

	l, r := int64(1), int64(minTime*totalTrips)

	for l < r {
		mid := (l + r) / 2
		sum := int64(0)
		for i := 0; i < len(time); i++ {
			sum += mid / int64(time[i])
		}

		if sum >= int64(totalTrips) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
