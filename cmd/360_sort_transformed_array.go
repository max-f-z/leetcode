package main

//lint:ignore U1000 unused
func sortTransformedArray(nums []int, a int, b int, c int) []int {
	res := make([]int, len(nums))
	switch {
	case a > 0:
		idx1, idx2 := 0, len(nums)-1
		idx := len(nums) - 1
		for idx >= 0 {
			tmp1 := a*nums[idx1]*nums[idx1] + b*nums[idx1] + c
			tmp2 := a*nums[idx2]*nums[idx2] + b*nums[idx2] + c
			if tmp1 > tmp2 {
				res[idx] = tmp1
				idx1++
			} else {
				res[idx] = tmp2
				idx2--
			}
			idx--
		}
	case a < 0:
		idx1, idx2 := 0, len(nums)-1
		idx := 0
		for idx < len(nums) {
			tmp1 := a*nums[idx1]*nums[idx1] + b*nums[idx1] + c
			tmp2 := a*nums[idx2]*nums[idx2] + b*nums[idx2] + c
			if tmp1 < tmp2 {
				res[idx] = tmp1
				idx1++
			} else {
				res[idx] = tmp2
				idx2--
			}
			idx++
		}
	case a == 0:
		switch {
		case b > 0:
			idx := 0
			for idx < len(nums) {
				res[idx] = b*nums[idx] + c
				idx++
			}
		case b < 0:
			idx := len(nums) - 1
			for idx >= 0 {
				res[idx] = b*nums[len(nums)-1-idx] + c
				idx--
			}
		case b == 0:
			idx := 0
			for idx < len(nums) {
				res[idx] = c
				idx++
			}
		}
	}
	return res
}
