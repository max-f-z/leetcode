package main

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	res := [][]int{}
	for i, v := range nums {
		subRes := permute(removeFromSlice43(nums, i))

		for _, w := range subRes {
			res = append(res, append(w, v))
		}
	}

	return res
}

func removeFromSlice43(nums []int, idx int) []int {
	res := []int{}
	if idx == 0 {
		res = append(res, nums[1:]...)
		return res
	}

	if idx == len(nums)-1 {
		res = append(res, nums[:len(nums)-1]...)
		return res
	}

	res = append(res, nums[:idx]...)
	res = append(res, nums[idx+1:]...)
	return res
}
