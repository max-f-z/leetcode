package main

import "container/heap"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type pq2583 []int64

func (pq pq2583) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq pq2583) Len() int {
	return len(pq)
}

func (pq pq2583) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pq2583) Push(x interface{}) {
	*pq = append(*pq, x.(int64))
}

func (pq *pq2583) Pop() (v interface{}) {
	v, *pq = (*pq)[pq.Len()-1], (*pq)[:(pq.Len()-1)]
	return v
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := []*TreeNode{root}

	mh := &pq2583{}
	heap.Init(mh)

	for len(queue) != 0 {
		var sum int64
		newQ := []*TreeNode{}
		for _, node := range queue {
			sum += int64(node.Val)
			if node.Left != nil {
				newQ = append(newQ, node.Left)
			}
			if node.Right != nil {
				newQ = append(newQ, node.Right)
			}
		}

		heap.Push(mh, sum)
		if mh.Len() > k {
			heap.Pop(mh)
		}
		queue = newQ
	}
	if mh.Len() < k {
		return -1
	}

	return heap.Pop(mh).(int64)
}
