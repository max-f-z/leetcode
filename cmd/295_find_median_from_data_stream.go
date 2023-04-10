package main

import "container/heap"

type minHeap []int

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() (v interface{}) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return v
}

func (h *minHeap) Top() int {
	return (*h)[0]
}

type maxHeap []int

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() (v interface{}) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return v
}

func (h *maxHeap) Top() int {
	return (*h)[0]
}

type MedianFinder struct {
	small *maxHeap
	big   *minHeap
}

/** initialize your data structure here. */
func Constructor295() MedianFinder {
	big := &minHeap{}
	small := &maxHeap{}
	heap.Init(big)
	heap.Init(small)
	mf := MedianFinder{
		small: small,
		big:   big,
	}

	return mf
}

//lint:ignore ST1006 this
func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.small, num)
	heap.Push(this.big, this.small.Top())
	heap.Pop(this.small)

	if this.small.Len() < this.big.Len() {
		heap.Push(this.small, this.big.Top())
		heap.Pop(this.big)
	}
}

//lint:ignore ST1006 this
func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.big.Len() {
		return float64(this.small.Top())
	}

	return float64(this.small.Top()+this.big.Top()) * 0.5
}

// o(n)
//
//lint:ignore ST1006 this
func (this *MedianFinder) Remove(num int) {

	found := false
	for i := 0; i < this.small.Len(); i++ {
		if (*this.small)[i] == num {
			heap.Remove(this.small, i)
			found = true
			break
		}
	}

	if !found {
		for i := 0; i < this.big.Len(); i++ {
			if (*this.big)[i] == num {
				heap.Remove(this.big, i)
				break
			}
		}
	}

	if this.small.Len() < this.big.Len() {
		heap.Push(this.small, this.big.Top())
		heap.Pop(this.big)
	}
}
