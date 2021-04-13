package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	mh := maxHeap(nums)
	heap.Init(&mh)

	val := 0

	for i := 0; i < k; i++ {
		val = heap.Pop(&mh).(int)
	}

	return val
}
