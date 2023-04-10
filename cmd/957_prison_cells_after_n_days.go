package main

//lint:ignore U1000 unused
func prisonAfterNDays(cells []int, n int) []int {
	var cur, original [8]int

	for i := 0; i < 8; i++ {
		cur[i] = cells[i]
	}

	for i := 0; i < n; i++ {
		tmp := [8]int{}

		for j := 1; j < 7; j++ {
			if cur[j-1] == cur[j+1] {
				tmp[j] = 1
			}
		}

		cur = tmp

		if i == 0 {
			original = cur
		} else if cur == original {
			i = (n - 1) / i * i
		}

	}

	for i := 0; i < 8; i++ {
		cells[i] = cur[i]
	}

	return cells
}
