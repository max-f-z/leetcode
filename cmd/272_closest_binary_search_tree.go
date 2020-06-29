package main

import (
	"container/heap"
	"math"
)

func closestKValues(root *TreeNode, target float64, k int) []int {
	res := make([]int, k)

	minHeap := &Heap272{Target: target}
	heap.Init(minHeap)

	closestKValuesHelper(root, minHeap)

	for i := 0; i < k; i++ {
		res[i] = heap.Pop(minHeap).(int)
	}

	return res
}

func closestKValuesHelper(node *TreeNode, minHeap *Heap272) {
	if node == nil {
		return
	}
	heap.Push(minHeap, node.Val)
	closestKValuesHelper(node.Left, minHeap)
	closestKValuesHelper(node.Right, minHeap)
}

type Heap272 struct {
	Val    []int
	Target float64
}

func (h *Heap272) Less(i, j int) bool {
	abs1 := math.Abs(float64(h.Val[i]) - h.Target)
	abs2 := math.Abs(float64(h.Val[j]) - h.Target)
	return abs1 < abs2
}

func (h *Heap272) Swap(i, j int) {
	h.Val[i], h.Val[j] = h.Val[j], h.Val[i]
}

func (h *Heap272) Len() int {
	return len(h.Val)
}

func (h *Heap272) Push(x interface{}) {
	h.Val = append(h.Val, x.(int))
}

func (h *Heap272) Pop() (v interface{}) {
	h.Val, v = h.Val[:len(h.Val)-1], h.Val[len(h.Val)-1]
	return v
}
