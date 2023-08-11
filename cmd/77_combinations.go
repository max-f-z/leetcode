package main

//lint:ignore U1000 unused
func combine(n int, k int) [][]int {
	coms := [][]int{}

	combineBacktracking(1, n, k, []int{}, &coms)

	return coms
}

func combineBacktracking(start int, n int, k int, curr []int, combinations *[][]int) {
	if len(curr) == k {
		res := make([]int, k)
		copy(res, curr)
		*combinations = append(*combinations, res)
		return
	}

	for i := start; i <= n; i++ {
		curr = append(curr, i)
		combineBacktracking(i+1, n, k, curr, combinations)
		curr = curr[:len(curr)-1]
	}
}
