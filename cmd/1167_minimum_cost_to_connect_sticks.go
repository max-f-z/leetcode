package main

import "container/heap"

func connectSticks(sticks []int) int {
	if len(sticks) == 1 {
		return 0
	}

	mh := append(minHeap{}, sticks...)
	heap.Init(&mh)

	cost := 0
	for mh.Len() != 1 {
		c := heap.Pop(&mh).(int) + heap.Pop(&mh).(int)
		cost += c

		heap.Push(&mh, c)
	}

	return cost
}
