package main

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			val := nums[i] + nums[j] + nums[k]

			switch {
			case val == 0:
				res = append(res, []int{nums[i], nums[j], nums[k]})
				tmp := j
				for tmp < k && nums[tmp] == nums[j] {
					tmp++
				}
				j = tmp
			case val > 0:
				k--
			case val < 0:
				j++
			default:
				k--
			}
		}
	}

	return res
}
