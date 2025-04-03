package main

func maximumTripletValue(nums []int) int64 {
	leftMax := make([]int, len(nums))
	rightMax := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		leftMax[i] = max(leftMax[i-1], nums[i-1])
		rightMax[len(nums)-1-i] = max(rightMax[len(nums)-i], nums[len(nums)-i])
	}

	maxValue := int64(0)
	for i := 1; i < len(nums)-1; i++ {
		value := int64(leftMax[i]-nums[i]) * int64(rightMax[i])
		maxValue = max(maxValue, value)
	}
	return maxValue
}
