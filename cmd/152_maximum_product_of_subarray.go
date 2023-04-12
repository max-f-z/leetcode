package main

//lint:ignore U1000 unused
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max_local := nums[0]
	min_local := nums[0]

	result := max_local

	for i := 1; i < len(nums); i++ {
		bigger, smaller := compare(max_local*nums[i], min_local*nums[i])

		max_local, _ = compare(bigger, nums[i])
		_, min_local = compare(smaller, nums[i])

		result, _ = compare(max_local, result)
	}

	return result

}

func compare(a, b int) (bigger, smaller int) {
	if a < b {
		return b, a
	}
	return a, b
}
