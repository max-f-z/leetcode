package main

import "container/heap"

/*
https://www.youtube.com/watch?v=5xT5GMTFvRI
*/
func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQ, 0)

	heap.Init(&pq)

	for _, v := range lists {
		if v != nil {
			heap.Push(&pq, v)
			v = v.Next
		}
	}

	dummy := &ListNode{}
	cur := dummy

	for pq.Len() > 0 {
		cur.Next = heap.Pop(&pq).(*ListNode)
		cur = cur.Next
		if cur.Next != nil {
			heap.Push(&pq, cur.Next)
		}
	}

	return dummy.Next
}

type PQ []*ListNode

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PQ) Pop() (v interface{}) {
	*pq, v = (*pq)[:pq.Len()-1], (*pq)[pq.Len()-1]
	return v
}
