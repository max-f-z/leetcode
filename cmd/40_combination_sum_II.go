package main

import "sort"

//lint:ignore U1000 unused
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSum2Helper(candidates, target)
}

func combinationSum2Helper(candidates []int, target int) [][]int {
	res := [][]int{}
	tmp := 0
	for i, v := range candidates {

		// 关键步骤 去重
		if tmp == v {
			continue
		}
		tmp = v

		newT := target - v

		if newT == 0 {
			return append(res, []int{v})
		}

		if newT < 0 {
			continue
		}

		// candidates[i:] 关键步骤 去重
		subRes := combinationSum2Helper(candidates[(i+1):], newT)

		for _, sr := range subRes {
			res = append(res, append(sr, v))
		}
	}

	return res
}
