package main

func rearrangeArray(nums []int) []int {
	p1, p2 := 0, 1

	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans[p1] = nums[i]
			p1 += 2
		} else {
			ans[p2] = nums[i]
			p2 += 2
		}
	}

	return ans
}
