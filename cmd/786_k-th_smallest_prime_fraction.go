package main

import "container/heap"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	pq := &pqpair{}
	heap.Init(pq)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			pair := &FractionPair{
				Fraction: float64(arr[i]) / float64(arr[j]),
				Num:      arr[i],
				Den:      arr[j],
			}

			heap.Push(pq, pair)
		}
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(pq)
	}

	target := heap.Pop(pq).(*FractionPair)
	return []int{target.Num, target.Den}
}

type FractionPair struct {
	Fraction float64
	Num      int
	Den      int
}

type pqpair []*FractionPair

func (pq pqpair) Len() int {
	return len(pq)
}

// make sure the pq top in queue is smallest
func (pq pqpair) Less(i, j int) bool {
	return pq[i].Fraction < pq[j].Fraction
}

func (pq pqpair) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqpair) Push(x interface{}) {
	*pq = append(*pq, x.(*FractionPair))
}

func (pq *pqpair) Pop() (v interface{}) {
	*pq, v = (*pq)[:pq.Len()-1], (*pq)[pq.Len()-1]
	return v
}
