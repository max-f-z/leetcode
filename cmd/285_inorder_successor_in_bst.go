package main

import "container/heap"

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	minHeap := &Heap285{}
	heap.Init(minHeap)

	inorderSuccessorHelper(root, minHeap, p)

	if minHeap.Len() > 0 {
		return heap.Pop(minHeap).(*TreeNode)
	}
	return nil
}

func inorderSuccessorHelper(cur *TreeNode, minHeap *Heap285, p *TreeNode) {
	if cur == nil {
		return
	}

	if cur.Val > p.Val {
		heap.Push(minHeap, cur)
	}
	inorderSuccessorHelper(cur.Left, minHeap, p)
	inorderSuccessorHelper(cur.Right, minHeap, p)
}

type Heap285 []*TreeNode

func (h Heap285) Len() int {
	return len(h)
}

func (h Heap285) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap285) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *Heap285) Push(x interface{}) {
	*h = append(*h, x.(*TreeNode))
}

func (h *Heap285) Pop() (v interface{}) {
	v, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return v
}
