package main

import (
	"container/heap"
	"sort"
)

type maxHeapSlice [][]int

func (mh *maxHeapSlice) Len() int {
	return len(*mh)
}

func (mh *maxHeapSlice) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *maxHeapSlice) Less(i, j int) bool {
	return (*mh)[i][0] > (*mh)[j][0]
}

func (mh *maxHeapSlice) Push(x interface{}) {
	*mh = append(*mh, x.([]int))
}

func (mh *maxHeapSlice) Pop() (v interface{}) {
	v, *mh = (*mh)[len(*mh)-1], (*mh)[:len(*mh)-1]
	return v
}

func (mh *maxHeapSlice) Top() []int {
	return (*mh)[0]
}

func findClosestElements(arr []int, k int, x int) []int {

	mh := &maxHeapSlice{}
	heap.Init(mh)

	for i := 0; i < len(arr); i++ {
		val := arr[i] - x
		if val < 0 {
			val = val * -1
		}

		if mh.Len() < k {
			heap.Push(mh, []int{val, arr[i]})
			continue
		}

		if val < mh.Top()[0] {
			heap.Pop(mh)
			heap.Push(mh, []int{val, arr[i]})
		}
	}

	res := []int{}

	for mh.Len() > 0 {
		tmp := mh.Top()[1]
		res = append(res, tmp)
		heap.Pop(mh)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}
