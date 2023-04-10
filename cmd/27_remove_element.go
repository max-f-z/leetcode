package main

//lint:ignore U1000 unused
func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return len(nums)
	}

	i, j := 0, 0
	for j < len(nums) {
		if nums[j] == val {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}

	return i
}
