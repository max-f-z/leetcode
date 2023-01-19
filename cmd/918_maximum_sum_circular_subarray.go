package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	left, right := make([]int, len(nums)), make([]int, len(nums))
	lMax, rMax := math.MinInt, math.MinInt
	lSum, rSum := 0, 0

	for i := 0; i < len(nums); i++ {
		lSum += nums[i]
		rSum += nums[len(nums)-i-1]

		lMax = maxInt(lSum, lMax)
		left[i] = lMax

		rMax = maxInt(rSum, rMax)
		right[len(nums)-i-1] = rMax
	}

	max := math.MinInt

	for i := 0; i < len(nums)-1; i++ {
		if max < left[i]+right[i+1] {
			max = left[i] + right[i+1]
		}
	}

	normalMax, localMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		localMax = maxInt(localMax+nums[i], nums[i])
		normalMax = maxInt(normalMax, localMax)
	}

	return maxInt(max, normalMax)
}
