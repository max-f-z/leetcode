package main

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	cnt := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}

	sort.Ints(nums)

	splitCnt := 0

	for i := 0; i < len(nums); i++ {
		if cnt[nums[i]] > 0 {
			cnt[nums[i]]--

			for j := 1; j < k; j++ {
				if cnt[nums[i]+j] == 0 {
					return false
				}
				cnt[nums[i]+j]--
			}
			splitCnt++
		}
	}

	return splitCnt*k == len(nums)
}
