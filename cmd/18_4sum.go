package main

import "sort"

//lint:ignore U1000 unused
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	if len(nums) < 4 {
		return [][]int{}
	}

	res := [][]int{}

	for l := 0; l < len(nums)-3; l++ {
		for ll := l + 1; ll < len(nums)-2; ll++ {
			h, hh := ll+1, len(nums)-1
			for h < hh {
				val := nums[l] + nums[ll] + nums[h] + nums[hh]
				switch {
				case val == target:
					res = append(res, []int{nums[l], nums[ll], nums[h], nums[hh]})
					for h < hh && nums[h] == nums[h+1] {
						h++
					}
					h++

					for h < hh && nums[hh] == nums[hh-1] {
						hh--
					}
					hh--
				case val < target:
					h++
				case val > target:
					hh--
				default:
					hh--
				}
			}

			for ll < len(nums)-2 && nums[ll] == nums[ll+1] {
				ll++
			}
		}

		for l < len(nums)-3 && nums[l] == nums[l+1] {
			l++
		}
	}
	return res
}
