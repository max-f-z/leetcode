package main

func bestClosingTime(customers string) int {
	n := len(customers)
	prefixY, prefixN := make([]int, n), make([]int, n)
	sumY, sumN := 0, 0

	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			sumY++
		} else {
			sumN++
		}

		prefixY[i] = sumY
		prefixN[i] = sumN
	}

	optimal := n + 1
	res := -1

	for i := 0; i < n; i++ {
		val := prefixN[i] + (sumY - prefixY[i])

		if customers[i] == 'Y' {
			val++
		} else if prefixN[i] > 0 {
			val--
		}

		if val < optimal {
			optimal = val
			res = i
		}
	}

	if optimal > sumN {
		res = n
	}

	return res
}
