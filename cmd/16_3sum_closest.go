package main

import (
	"math"
	"sort"
)

//lint:ignore U1000 unused
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	min := math.MaxInt32
	res := 0

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			val := nums[i] + nums[j] + nums[k]

			switch {
			case val == target:
				return target
			case val > target:
				tmp := int(math.Abs(float64(val - target)))
				if tmp < min {
					min = tmp
					res = val
				}
				k--
			case val < target:
				tmp := int(math.Abs(float64(val - target)))
				if tmp < min {
					min = tmp
					res = val
				}
				j++
			}
		}
	}
	return res
}
