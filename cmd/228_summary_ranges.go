package main

import (
	"fmt"
	"strconv"
)

//lint:ignore U1000 unused
func summaryRanges(nums []int) []string {
	res := []string{}

	if len(nums) == 0 {
		return res
	}

	if len(nums) == 1 {
		return append(res, strconv.Itoa(nums[0]))
	}

	prefix := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if prefix[len(prefix)-1]+1 == nums[i] {
			prefix = append(prefix, nums[i])
		} else {
			if len(prefix) == 1 {
				res = append(res, strconv.Itoa(prefix[0]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", prefix[0], prefix[len(prefix)-1]))
			}
			prefix = []int{nums[i]}
		}
	}

	if len(prefix) == 1 {
		res = append(res, strconv.Itoa(prefix[0]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", prefix[0], prefix[len(prefix)-1]))
	}

	return res
}
