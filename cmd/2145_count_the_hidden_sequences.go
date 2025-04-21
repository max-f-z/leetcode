package main

func numberOfArrays(differences []int, lower int, upper int) int {
	n := len(differences)

	offsets := make([]int, n)

	var low, high int
	if differences[0] > 0 {
		low, high = 0, differences[0]
	} else {
		low, high = differences[0], 0
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += differences[i]
		offsets[i] = sum

		if offsets[i] < low {
			low = offsets[i]
		}

		if offsets[i] > high {
			high = offsets[i]
		}
	}

	minOffset := lower - low
	high += minOffset

	ans := upper - high + 1

	if ans < 0 {
		return 0
	}
	return ans
}
