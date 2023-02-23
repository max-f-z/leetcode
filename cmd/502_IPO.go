package main

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	l := len(profits)

	ps := []*project{}
	for i := 0; i < l; i++ {
		p := &project{
			profit:  profits[i],
			capital: capital[i],
		}

		ps = append(ps, p)
	}

	pss := projects(ps)
	sort.Sort(pss)

	idx := 0

	maxHeap := &pq{}
	heap.Init(maxHeap)

	for i := 0; i < k; i++ {
		for ; idx < l && pss[idx].capital <= w; idx++ {
			heap.Push(maxHeap, pss[idx])
		}

		if maxHeap.Len() > 0 {
			selected := heap.Pop(maxHeap)
			w += (selected.(*project)).profit
		}
	}

	return w
}

type project struct {
	profit  int
	capital int
}

type projects []*project

func (p projects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p projects) Less(i, j int) bool {
	return p[i].capital < p[j].capital
}

func (p projects) Len() int {
	return len(p)
}

type pq []*project

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pq) Less(i, j int) bool {
	// maxHeap
	return p[i].profit > p[j].profit
}

func (p pq) Len() int {
	return len(p)
}

func (p *pq) Push(x interface{}) {
	(*p) = append(*p, x.(*project))
}

func (p *pq) Pop() (v interface{}) {
	(*p), v = (*p)[:p.Len()-1], (*p)[p.Len()-1]
	return v
}
