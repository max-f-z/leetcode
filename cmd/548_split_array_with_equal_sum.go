package main

func splitArray548(nums []int) bool {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i-1]
	}

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				if sums[i]-nums[i] == sums[j-1]-sums[i] && sums[i]-nums[i] == sums[k-1]-sums[j] && sums[i]-nums[i] == sums[len(nums)-1]-sums[k] {
					return true
				}
			}
		}
	}

	return false
}
