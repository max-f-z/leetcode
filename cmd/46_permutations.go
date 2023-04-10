package main

//lint:ignore U1000 unused
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

type permuteType struct {
	res   [][]int
	usage map[int]int
	nums  []int
}

//lint:ignore U1000 unused
func permuteRefactor(nums []int) [][]int {
	usage := map[int]int{}

	for i := 0; i < len(nums); i++ {
		usage[nums[i]] = 0
	}

	pt := &permuteType{
		res:   [][]int{},
		nums:  nums,
		usage: usage,
	}

	pt.permuteHelper([]int{})

	return pt.res
}

func (pt *permuteType) permuteHelper(prefix []int) {
	if len(prefix) == len(pt.nums) {
		ans := make([]int, len(pt.nums))
		copy(ans, prefix)
		pt.res = append(pt.res, ans)
		return
	}

	for k, v := range pt.usage {
		if v == 0 {
			pt.usage[k] = 1
			prefix = append(prefix, k)
			pt.permuteHelper(prefix)
			pt.usage[k] = 0
			prefix = prefix[:len(prefix)-1]
		}
	}
}
