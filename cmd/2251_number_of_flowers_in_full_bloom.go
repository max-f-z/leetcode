package main

import (
	"container/heap"
	"sort"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	sort.Slice(flowers, func(i, j int) bool {
		a, b := flowers[i], flowers[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})

	ppl := make([]int, len(people))
	copy(ppl, people)
	sort.Ints(ppl)

	res := make([]int, len(ppl))

	ans := map[int]int{}

	mh := &minHeap{}
	heap.Init(mh)

	idx := 0

	for _, v := range ppl {
		for idx < len(flowers) && flowers[idx][0] <= v {
			heap.Push(mh, flowers[idx][1])
			idx++
		}
		for mh.Len() > 0 && mh.Top() < v {
			heap.Pop(mh)
		}

		ans[v] = mh.Len()
	}

	for i, v := range people {
		res[i] = ans[v]
	}
	return res
}
