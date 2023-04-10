package main

import "math"

//lint:ignore U1000 unused
func maxProfit(prices []int) int {
	max := 0

	min := math.MaxInt32

	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}

		if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}
