package main

import (
	"fmt"
)

func main() {

	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			continue
		}

		if nums[i] == target {
			return i
		}

		if nums[i] > target {
			return i
		}
	}

	return len(nums)
}
