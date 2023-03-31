package main

type subset struct {
	res  [][]int
	nums []int
}

func subsets(nums []int) [][]int {
	s := &subset{
		res:  [][]int{},
		nums: nums,
	}

	s.subsetshelper(0, []int{})

	return s.res
}

func (s *subset) subsetshelper(start int, path []int) {
	resPath := make([]int, len(path))
	copy(resPath, path)
	s.res = append(s.res, resPath)

	for i := start; i < len(s.nums); i++ {
		path = append(path, s.nums[i])

		s.subsetshelper(i+1, path)

		path = path[:len(path)-1]
	}
}
