package main

func minOperations3375(nums []int, k int) int {
	cnt := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			return -1
		} else if nums[i] > k {
			cnt[nums[i]] = true
		}
	}

	return len(cnt)
}
