package main

func subarraySum(nums []int, k int) int {
	sum := 0
	num := 0

	tmp := map[int]int{}
	tmp[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if tmp[sum-k] != 0 {
			num += tmp[sum-k]
		}

		if tmp[sum] == 0 {
			tmp[sum] = 1
		} else {
			tmp[sum] += 1
		}
	}
	return num
}
