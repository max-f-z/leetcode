package main

func tupleSameProduct(nums []int) int {
	ref := map[int]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			ref[nums[i]*nums[j]] += 1
		}
	}

	pairs := 0

	for _, v := range ref {
		pairs += (v - 1) * v / 2
	}

	return pairs * 8
}
