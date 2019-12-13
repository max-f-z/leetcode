package main

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return Max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}
