package main

import (
	"math"
)

//lint:ignore U1000 unused
func divide(dividend int, divisor int) int {
	res := 0

	switch divisor {
	case 1:
		res = dividend
	case -1:
		res = -dividend
	default:
		a := math.Abs(float64(dividend))
		b := math.Abs(float64(divisor))

		sum := b
		for sum <= a {
			sum += b
			res++
		}
	}

	if res >= math.MaxInt32 {
		return math.MaxInt32
	}

	if res <= math.MinInt32 {
		return math.MinInt32
	}

	if divisor != 1 && divisor != -1 && math.Abs(float64(dividend+divisor)) != math.Abs(float64(dividend))+math.Abs(float64(divisor)) {
		return -res
	}

	return res
}
