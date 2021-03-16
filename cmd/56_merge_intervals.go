package main

import (
	"container/heap"
	"sort"
)

func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})

	res := [][]int{}

	mh := &minHeap{}
	heap.Init(mh)

	start := -1

	for idx, v := range intervals {
		if idx == 0 {
			start = v[0]
			heap.Push(mh, v[1])
			continue
		}

		if mh.Top() < v[0] {
			res = append(res, []int{start, heap.Pop(mh).(int)})
			start = v[0]
			heap.Push(mh, v[1])
			continue
		}

		if mh.Top() > v[1] {
			continue
		}
		heap.Pop(mh)
		heap.Push(mh, v[1])
	}

	if mh.Len() > 0 {
		res = append(res, []int{start, heap.Pop(mh).(int)})
	}

	return res
}
