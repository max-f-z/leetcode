package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSumHelper(candidates, target)
}

func combinationSumHelper(candidates []int, target int) [][]int {
	res := [][]int{}

	for i, v := range candidates {
		newT := target - v

		if newT == 0 {
			return append(res, []int{v})
		}

		if newT < 0 {
			continue
		}

		subRes := combinationSum(candidates[i:], newT)

		for _, sr := range subRes {
			res = append(res, append(sr, v))
		}
	}

	return res
}
