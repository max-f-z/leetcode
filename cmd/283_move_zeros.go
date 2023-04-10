package main

//lint:ignore U1000 unused
func moveZeroes(nums []int) {
	idx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}
}
