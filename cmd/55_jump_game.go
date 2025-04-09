package main

func canJump(nums []int) bool {
	mem := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		mem[i] = -1
	}

	var process func([]int, []int, int) int
	process = func(nums []int, mem []int, current int) int {
		if mem[current] != -1 {
			return mem[current]
		}

		if current == len(nums)-1 {
			return 1
		}

		for i := 1; i <= nums[current]; i++ {
			result := process(nums, mem, current+i)
			if result == 1 {
				mem[current] = 1
				return 1
			}
		}
		mem[current] = 0
		return 0
	}

	return process(nums, mem, 0) == 1
}
