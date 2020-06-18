package main

func searchRange(nums []int, target int) []int {
	start := -1
	end := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target && start == -1 {
			start = i
			end = i
			continue
		}

		if nums[i] == target && start != -1 {
			end = i
			continue
		}
	}

	return []int{start, end}
}
