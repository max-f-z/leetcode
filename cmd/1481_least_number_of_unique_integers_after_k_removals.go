package main

import "container/heap"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnt := map[int]int{}

	for i := 0; i < len(arr); i++ {
		cnt[arr[i]]++
	}

	freq := make([]int, len(cnt))
	idx := 0
	for _, v := range cnt {
		freq[idx] = v
		idx++
	}

	pq := pq1481(freq)
	heap.Init(&pq)

	for k > 0 {
		val := heap.Pop(&pq).(int)
		if k < val {
			return pq.Len() + 1
		}
		k -= val
	}

	return pq.Len()
}

type pq1481 []int

func (pq pq1481) Len() int {
	return len(pq)
}

func (pq pq1481) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq pq1481) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pq1481) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *pq1481) Pop() (v interface{}) {
	v, *pq = (*pq)[pq.Len()-1], (*pq)[:pq.Len()-1]
	return v
}
