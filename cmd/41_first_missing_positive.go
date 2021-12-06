package main

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// For numbers bigger then n or zero or negative. mark them to be n+1
	for i := 0; i < n; i++ {
		if nums[i] > n || nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		val := nums[i]
		if val < 0 {
			val *= -1
		}

		if (val-1) < n && nums[val-1] > 0 {
			nums[val-1] *= -1
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}
