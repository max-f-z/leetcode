package main

import "container/heap"

func minOperations3066(nums []int, k int) int {
	mh := minHeap(nums)
	heap.Init(&mh)

	step := 0

	for mh.Top() < k {
		op1 := heap.Pop(&mh).(int)
		op2 := heap.Pop(&mh).(int)

		res := min(op1, op2)*2 + max(op1, op2)

		heap.Push(&mh, res)
		step++
	}

	return step
}
