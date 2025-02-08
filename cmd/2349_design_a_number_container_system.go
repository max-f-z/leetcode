package main

import "container/heap"

type NumberContainers struct {
	indexToNumber map[int]int
	numberToIndex map[int]*minHeap
}

func Constructor2349() NumberContainers {
	nc := NumberContainers{
		indexToNumber: map[int]int{},
		numberToIndex: map[int]*minHeap{},
	}

	return nc
}

//lint:ignore U1000 unused
func (nc *NumberContainers) Change(index int, number int) {
	nc.indexToNumber[index] = number

	if mh, ok := nc.numberToIndex[number]; ok {
		heap.Push(mh, index)
	} else {
		mh := &minHeap{}

		heap.Init(mh)
		heap.Push(mh, index)

		nc.numberToIndex[number] = mh
	}
}

//lint:ignore U1000 unused
func (nc *NumberContainers) Find(number int) int {
	if nc.numberToIndex[number] == nil {
		return -1
	}

	mh := nc.numberToIndex[number]

	for mh.Len() > 0 {
		idx := mh.Top()
		if nc.indexToNumber[idx] == number {
			return idx
		}

		heap.Pop(mh)
	}

	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
