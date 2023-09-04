package main

func minimumReplacement(nums []int) int64 {
	var ans int64 = 0

	n := len(nums)

	currentMax := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] <= currentMax {
			currentMax = nums[i]
			continue
		}

		numElms := -1
		if nums[i]%currentMax == 0 {
			numElms = nums[i] / currentMax
		} else {
			numElms = nums[i]/currentMax + 1
		}

		ans += int64(numElms - 1)
		currentMax = nums[i] / numElms
	}

	return ans
}
