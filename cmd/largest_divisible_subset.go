package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	sort.Ints(nums)

	dpCnt := make([]int, len(nums))
	dpParent := make([]int, len(nums))

	maxLen := 0
	maxNum := 0

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 && dpCnt[i] < dpCnt[j]+1 {
				dpCnt[i] = dpCnt[j] + 1
				dpParent[i] = j
				if dpCnt[i] > maxLen {
					maxLen = dpCnt[i]
					maxNum = i
				}
			}
		}
	}

	ret := []int{}

	for i := 0; i < maxLen; i++ {
		ret = append(ret, nums[maxNum])
		maxNum = dpParent[maxNum]
	}

	return ret
}
