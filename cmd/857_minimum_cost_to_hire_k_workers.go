package main

import (
	"container/heap"
	"math"
	"sort"
)

type RatioQualityPair struct {
	Ratio   float64
	Quality int
}

type maxHeap857 []int

func (mh maxHeap857) Len() int {
	return len(mh)
}

func (mh maxHeap857) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh maxHeap857) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *maxHeap857) Push(x interface{}) {
	(*mh) = append(*mh, x.(int))
}

func (mh *maxHeap857) Pop() (v interface{}) {
	(*mh), v = (*mh)[:mh.Len()-1], (*mh)[mh.Len()-1]
	return v
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]*RatioQualityPair, len(quality))
	for i := 0; i < len(quality); i++ {
		workers[i] = &RatioQualityPair{
			Quality: quality[i],
			Ratio:   float64(wage[i]) / float64(quality[i]),
		}
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i].Ratio < workers[j].Ratio
	})

	ans := math.MaxFloat64
	totalQuality := float64(0)

	group := &maxHeap857{}
	heap.Init(group)

	for _, worker := range workers {
		heap.Push(group, worker.Quality)
		totalQuality += float64(worker.Quality)
		if group.Len() == k {
			val := worker.Ratio * totalQuality
			if ans > val {
				ans = val
			}
		} else if group.Len() > k {
			maxQuality := heap.Pop(group).(int)
			totalQuality -= float64(maxQuality)

			val := worker.Ratio * totalQuality
			if ans > val {
				ans = val
			}
		}
	}

	return ans
}
