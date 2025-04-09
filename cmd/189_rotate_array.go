package main

func rotate189(nums []int, k int) {
	k = k % len(nums)

	reverse := func(nums []int, start int, end int) {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
