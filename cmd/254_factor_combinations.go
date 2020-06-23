package main

import (
	"math"
)

func getFactors(n int) [][]int {
	if n <= 1 {
		return [][]int{}
	}

	res := [][]int{}

	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			res = append(res, []int{i, n / i})
			inner := getFactors(n / i)

			for j := 0; j < len(inner); j++ {
				if inner[j][0] >= i {
					inner[j] = append([]int{i}, inner[j]...)
					res = append(res, inner[j])
				}
			}
		}
	}

	return res
}
