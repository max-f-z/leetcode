package main

func longestConsecutiveII(nums []int) int {
	ref := map[int]int{}

	ans := 0

	for i := 0; i < len(nums); i++ {
		ref[nums[i]] = 1
	}

	for i := 0; i < len(nums); i++ {
		if ref[nums[i]-1] == 1 {
			continue
		}

		streak := 0
		cur := nums[i]
		for ref[cur] != 0 {
			streak++
			cur++
		}

		if streak > ans {
			ans = streak
		}
	}

	return ans
}
