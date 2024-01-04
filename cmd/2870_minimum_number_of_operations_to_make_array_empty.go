package main

func minOperations2870(nums []int) int {
	cnt := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}

	ans := 0
	for _, v := range cnt {
		if v == 1 {
			return -1
		} else {
			if v%3 == 0 {
				ans += v / 3
			} else {
				ans += v/3 + 1
			}
		}
	}

	return ans
}
