package main

import (
	"container/heap"
	"fmt"
)

func kClosest(points [][]int, k int) [][]int {
	mh := &maxHeapSlice{}

	heap.Init(mh)

	for i := 0; i < len(points); i++ {
		tmp := points[i][0]*points[i][0] + points[i][1]*points[i][1]

		if mh.Len() < k {
			heap.Push(mh, []int{tmp, i})
		} else {
			fmt.Println(mh.Top()[0])
			if mh.Top()[0] > tmp {
				heap.Pop(mh)
				heap.Push(mh, []int{tmp, i})
			}
		}
	}

	res := [][]int{}

	for i := 0; i < k; i++ {
		v := heap.Pop(mh).([]int)
		res = append(res, points[v[1]])
	}

	return res
}
