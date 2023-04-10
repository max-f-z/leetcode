package main

import "container/heap"

/*
https://www.youtube.com/watch?v=5xT5GMTFvRI
*/
//lint:ignore U1000 unused
func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQ, 0)

	heap.Init(&pq)

	for _, v := range lists {
		if v != nil {
			heap.Push(&pq, v)
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

type node struct {
	val int
	idx int
	pos int
}

type mhNode []*node

func (mh mhNode) Len() int {
	return len(mh)
}

func (mh mhNode) Less(i, j int) bool {
	return mh[i].val < mh[j].val
}

func (mh mhNode) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *mhNode) Push(x interface{}) {
	*mh = append(*mh, x.(*node))
}

func (mh *mhNode) Pop() (v interface{}) {
	*mh, v = (*mh)[:mh.Len()-1], (*mh)[mh.Len()-1]
	return v
}

//lint:ignore U1000 unused
func mergeKArrays(lists [][]int) []int {
	mh := &mhNode{}

	heap.Init(mh)

	for k, v := range lists {
		if len(v) > 0 {
			n := &node{
				val: v[0],
				idx: k,
				pos: 0,
			}

			heap.Push(mh, n)
		}
	}

	res := []int{}

	for mh.Len() > 0 {
		n := heap.Pop(mh).(*node)
		res = append(res, n.val)
		if n.pos+1 <= len(lists[n.idx])-1 {
			nn := &node{
				val: lists[n.idx][n.pos+1],
				idx: n.idx,
				pos: n.pos + 1,
			}
			heap.Push(mh, nn)
		}
	}

	return res
}
