package main

import "sort"

//lint:ignore U1000 unused
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	s := &subset{
		res:  [][]int{},
		nums: nums,
	}

	s.subsetshelperII(0, []int{})

	return s.res
}

func (s *subset) subsetshelperII(start int, path []int) {
	resPath := make([]int, len(path))
	copy(resPath, path)
	s.res = append(s.res, resPath)

	for i := start; i < len(s.nums); i++ {
		if i > start && s.nums[i] == s.nums[i-1] {
			continue
		}
		path = append(path, s.nums[i])

		s.subsetshelper(i+1, path)

		path = path[:len(path)-1]
	}
}
