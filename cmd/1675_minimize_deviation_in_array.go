package main

import "container/heap"

func minimumDeviation(nums []int) int {
	maxH := &maxHeap{}

	heap.Init(maxH)

	min := 0

	for i := 0; i < len(nums); i++ {
		var val int
		if nums[i]%2 == 0 {
			val = nums[i]
		} else {
			val = nums[i] * 2
		}

		heap.Push(maxH, val)

		if min == 0 || val < min {
			min = val
		}
	}

	max := maxH.Top()

	ans := max - min

	for max%2 == 0 {
		val := (heap.Pop(maxH)).(int)
		val = val / 2
		if val < min {
			min = val
		}
		heap.Push(maxH, val)
		max = maxH.Top()
		if max-min < ans {
			ans = max - min
		}
	}

	return ans
}
