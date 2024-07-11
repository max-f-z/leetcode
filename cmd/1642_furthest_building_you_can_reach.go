package main

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {

	costs := make([]int, len(heights)-1)

	for i := 1; i < len(heights); i++ {
		if heights[i] <= heights[i-1] {
			costs[i-1] = 0
		} else {
			costs[i-1] = heights[i] - heights[i-1]
		}
	}

	sum := 0
	idx := 0

	pq := &pq1642{}
	heap.Init(pq)

	for ; idx < len(costs); idx++ {
		heap.Push(pq, costs[idx])
		sum += costs[idx]

		if sum > bricks {
			if ladders > 0 {
				cost := heap.Pop(pq).(int)
				sum -= cost
				ladders--
				continue
			}
			break
		}
	}

	return idx
}

type pq1642 []int

func (pq pq1642) Len() int {
	return len(pq)
}

func (pq pq1642) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq pq1642) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *pq1642) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *pq1642) Pop() (v interface{}) {
	v, *pq = (*pq)[pq.Len()-1], (*pq)[:pq.Len()-1]
	return v
}
