package main

// sum(all) + m(n-1) = (a1+m) * n
// sum(all) + mn -m = a1 * n + mn
// m = sum(all) - a1 * n
//
//lint:ignore U1000 unused
func minMoves(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	sum := 0
	min := nums[0]
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if min > nums[i] {
			min = nums[i]
		}
	}

	return sum - min*n
}
