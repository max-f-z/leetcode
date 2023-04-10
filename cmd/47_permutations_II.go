package main

import "sort"

//lint:ignore U1000 unused
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return permuteUniqueHelper(nums)
}

func permuteUniqueHelper(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	res := [][]int{}
	tmp := -100
	for i, v := range nums {
		if tmp == v {
			continue
		}

		tmp = v

		subRes := permuteUniqueHelper(removeFromSlice47(nums, i))

		for _, w := range subRes {
			res = append(res, append(w, v))
		}
	}

	return res
}

func removeFromSlice47(nums []int, idx int) []int {
	res := []int{}
	if idx == 0 {
		res = append(res, nums[1:]...)
		return res
	}

	if idx == len(nums)-1 {
		res = append(res, nums[:len(nums)-1]...)
		return res
	}

	res = append(res, nums[:idx]...)
	res = append(res, nums[idx+1:]...)
	return res
}
