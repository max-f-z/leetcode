package main

func pivotIndex(nums []int) int {
	sum := 0
	prefixSum := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		prefixSum[i] = sum
	}

	for i := 0; i < len(nums); i++ {
		r := sum - prefixSum[i]
		l := prefixSum[i] - nums[i]

		if l == r {
			return i
		}
	}

	return -1
}
