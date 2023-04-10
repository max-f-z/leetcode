package main

//lint:ignore U1000 unused
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

func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start := -1
	end := -1

	low, up := 0, len(nums)-1
	for low < up {
		mid := (low + up) / 2

		if nums[mid] >= target {
			up = mid
		} else {
			low = mid + 1
		}
	}

	start = low

	low, up = start, len(nums)-1

	for low < up {
		mid := (low + up + 1) / 2
		if nums[mid] <= target {
			low = mid
		} else {
			up = mid - 1
		}
	}

	end = low

	return []int{start, end}
}
