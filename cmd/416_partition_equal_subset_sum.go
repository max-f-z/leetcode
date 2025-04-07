package main

func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	mem := make([][]int, n)

	for i := 0; i < n; i++ {
		mem[i] = make([]int, target+1)
		for j := 0; j < target+1; j++ {
			mem[i][j] = -1
		}
	}

	var dfs func(idx int, leftSum int) int

	dfs = func(idx, leftSum int) int {
		if leftSum == 0 {
			return 1
		}

		if idx == 0 || leftSum < 0 {
			return 0
		}

		if mem[idx][leftSum] != -1 {
			return mem[idx][leftSum]
		}

		mem[idx][leftSum] = dfs(idx-1, leftSum)
		result := dfs(idx-1, leftSum-nums[idx])
		if mem[idx][leftSum] == 0 {
			mem[idx][leftSum] = result
		}

		return mem[idx][leftSum]
	}

	return dfs(n-1, target) == 1
}
