package main

import (
	"container/heap"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	minHeap := &Heap{}
	heap.Init(minHeap)

	heap.Push(minHeap, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= minHeap.Top() {
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, intervals[i][1])
	}

	return minHeap.Len()
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h Heap) Top() int {
	return h[0]
}
