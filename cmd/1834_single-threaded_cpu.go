package main

import (
	"container/heap"
	"sort"
)

// sorting + minHeap
func getOrder(tasks [][]int) []int {
	t := cpuTasks([][]int{})
	for i := 0; i < len(tasks); i++ {
		t = append(t, []int{tasks[i][0], tasks[i][1], i})
	}

	sort.Sort(t)

	buf := &minHeap1834{}
	heap.Init(buf)

	ans := []int{}
	nextStart := 0

	for i := 0; i < len(tasks); i++ {
		if buf.Len() == 0 {
			heap.Push(buf, t[i])
			continue
		}
		if buf.Top()[0] == t[i][0] {
			heap.Push(buf, t[i])
			continue
		}
		if t[i][0] < nextStart {
			heap.Push(buf, t[i])
			continue
		}

		for t[i][0] > nextStart && buf.Len() > 0 {
			next := heap.Pop(buf).([]int)
			ans = append(ans, next[2])

			if next[0] > nextStart {
				nextStart = next[0]
			}
			nextStart = nextStart + next[1]
		}

		heap.Push(buf, t[i])
	}

	for buf.Len() > 0 {
		next := heap.Pop(buf).([]int)
		ans = append(ans, next[2])
		if next[0] > nextStart {
			nextStart = next[0]
		}
		nextStart = nextStart + next[1]
	}

	return ans
}

type minHeap1834 [][]int

func (mh minHeap1834) Less(i, j int) bool {
	if mh[i][1] == mh[j][1] {
		return mh[i][2] < mh[j][2]
	}
	return mh[i][1] < mh[j][1]
}

func (mh minHeap1834) Len() int {
	return len(mh)
}

func (mh minHeap1834) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap1834) Pop() (v interface{}) {
	v, *mh = (*mh)[len(*mh)-1], (*mh)[:len(*mh)-1]
	return v
}

func (mh *minHeap1834) Push(x interface{}) {
	*mh = append(*mh, x.([]int))
}

func (mh *minHeap1834) Top() []int {
	return (*mh)[0]
}

type cpuTasks [][]int

func (c cpuTasks) Less(i, j int) bool {
	return c[i][0] < c[j][0]
}

func (c cpuTasks) Len() int {
	return len(c)
}

func (c cpuTasks) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
