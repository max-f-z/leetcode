package main

type cs struct {
	mem  map[int]int
	nums []int
}

func (cm cs) getSum(target int) int {
	if target == 0 {
		return 1
	}
	if val, ok := cm.mem[target]; ok {
		return val
	}
	ans := 0
	for i := 0; i < len(cm.nums); i++ {
		if target-cm.nums[i] >= 0 {
			val := cm.getSum(target - cm.nums[i])
			ans += val
		}
	}
	cm.mem[target] = ans
	return ans
}

func combinationSum4(nums []int, target int) int {
	cm := cs{
		mem:  map[int]int{},
		nums: nums,
	}

	return cm.getSum(target)
}
