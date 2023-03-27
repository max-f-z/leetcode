package main

import "container/heap"

type pair struct {
	val int
	cnt int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	pq := make(PairPQ, 0)
	heap.Init(&pq)

	for k, v := range m {
		heap.Push(&pq, &pair{val: k, cnt: v})
	}

	ret := make([]int, 0)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(&pq).(*pair).val)
	}

	return ret
}

type PairPQ []*pair

func (pq PairPQ) Len() int {
	return len(pq)
}

func (pq PairPQ) Less(i, j int) bool {
	return pq[i].cnt > pq[j].cnt
}

func (pq PairPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PairPQ) Push(x interface{}) {
	(*pq) = append((*pq), x.(*pair))
}

func (pq *PairPQ) Pop() (v interface{}) {
	*pq, v = (*pq)[:pq.Len()-1], (*pq)[pq.Len()-1]
	return v
}
