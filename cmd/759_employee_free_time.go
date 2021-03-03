package main

import (
	"container/heap"
)

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	minHeap := &Heap759{}
	heap.Init(minHeap)

	for _, v := range schedule {
		for _, u := range v {
			heap.Push(minHeap, u)
		}
	}

	start := -1
	// end := math.MaxInt32

	res := []*Interval{}

	for minHeap.Top() != nil {
		s := heap.Pop(minHeap).(*Interval)
		if start == -1 {
			start = s.End
			continue
		}

		if s.Start > start {
			interval := &Interval{
				Start: start,
				End:   s.Start,
			}
			res = append(res, interval)
			start = s.End
			continue
		}

		if s.End > start {
			start = s.End
			continue
		}

	}

	return res
}

type Heap759 []*Interval

func (h Heap759) Len() int {
	return len(h)
}

func (h Heap759) Less(i, j int) bool {
	if h[i].Start != h[j].Start {
		return h[i].Start < h[j].Start
	}
	return h[i].End < h[j].End
}

func (h Heap759) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap759) Push(x interface{}) {
	*h = append(*h, x.(*Interval))
}

func (h *Heap759) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h Heap759) Top() *Interval {
	if h.Len() == 0 {
		return nil
	}
	return h[0]
}
