package main

import "math"

//lint:ignore U1000 unused
func reverse(x int) int {
	res := 0
	for {
		if x == 0 {
			break
		}
		res = res*10 + (x % 10)
		x = x / 10

		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}

	return res
}
