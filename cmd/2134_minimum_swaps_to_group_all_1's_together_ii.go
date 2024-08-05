package main

func minSwaps(nums []int) int {
	dup := append(nums, nums...)

	totalOnes := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			totalOnes++
		}
	}

	numOfOnes := 0
	for i := 0; i < totalOnes; i++ {
		if dup[i] == 1 {
			numOfOnes++
		}
	}

	minSwap := totalOnes - numOfOnes
	for i := 1; i < len(nums); i++ {
		numOfOnes -= dup[i-1]
		numOfOnes += dup[i+totalOnes-1]

		minSwap = min(minSwap, totalOnes-numOfOnes)
	}

	return minSwap
}
