package main

func maxSubArray(nums []int) int {
	max := nums[0]
	current := nums[0]

	for i := 1; i < len(nums); i++ {
		if current+nums[i] > nums[i] {
			current = current + nums[i]
		} else {
			current = nums[i]
		}

		if current > max {
			max = current
		}
	}

	return max
}
