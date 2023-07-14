package main

func movesToMakeZigzag(nums []int) int {
	n := len(nums)
	even := 0
	odd := 0

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			// Even index
			if i > 0 && nums[i] >= nums[i-1] {
				even += nums[i] - nums[i-1] + 1
			}
			if i < n-1 && nums[i] >= nums[i+1] {
				even += nums[i] - nums[i+1] + 1
			}
		} else {
			// Odd index
			if i > 0 && nums[i] <= nums[i-1] {
				odd += nums[i-1] - nums[i] + 1
			}
			if i < n-1 && nums[i] <= nums[i+1] {
				odd += nums[i+1] - nums[i] + 1
			}
		}
	}

	if even < odd {
		return even
	}
	return odd
}
