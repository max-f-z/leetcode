package main

import "sort"

//lint:ignore U1000 unused
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

		// candidates[i:] 关键步骤 去重
		subRes := combinationSumHelper(candidates[i:], newT)

		for _, sr := range subRes {
			res = append(res, append(sr, v))
		}
	}

	return res
}

type combinationSumType struct {
	candidates []int
	min        int
	res        [][]int
}

func combinationSumII(candidates []int, target int) [][]int {
	min := 50

	for i := 0; i < len(candidates); i++ {
		if candidates[i] < min {
			min = candidates[i]
		}
	}

	cst := &combinationSumType{
		candidates: candidates,
		res:        [][]int{},
		min:        min,
	}

	cst.helper(target, []int{}, 0)

	return cst.res
}

func (cst *combinationSumType) helper(remain int, prefix []int, start int) {
	if remain == 0 {
		ans := make([]int, len(prefix))
		copy(ans, prefix)
		cst.res = append(cst.res, ans)
	}

	if remain < cst.min {
		return
	}

	for i := start; i < len(cst.candidates); i++ {
		prefix = append(prefix, cst.candidates[i])
		cst.helper(remain-cst.candidates[i], prefix, i)
		prefix = prefix[:len(prefix)-1]
	}
}
