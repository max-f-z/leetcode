package main

func minimumOperations3396(nums []int) int {
	cnt := map[int]int{}

	ans := 0

	for i := 0; i < len(nums); i++ {
		cnt[nums[i]]++
	}

	start_idx := 0

	for len(nums) >= 0 {
		flag := false
		for _, v := range cnt {
			if v > 1 {
				flag = true
				break
			}
		}
		if !flag {
			break
		}

		for i := 0; i < 3 && (start_idx+i) < len(nums); i++ {
			cnt[nums[start_idx+i]]--
		}
		start_idx += 3
		ans++
	}

	return ans
}
