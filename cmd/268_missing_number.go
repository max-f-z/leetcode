package main

func missingNumber(nums []int) int {
	expectedSum := len(nums) * (len(nums) + 1) / 2
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return expectedSum - sum
}
