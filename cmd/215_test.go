package main

import (
	"container/heap"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func Test215(t *testing.T) {
	mh := &maxHeap{3, 2, 1, 5, 6, 4}

	heap.Init(mh)

	l := mh.Len()

	for i := 0; i < l; i++ {
		tmp := heap.Pop(mh).(int)
		fmt.Println(tmp)
	}

	nums := []int{3, 2, 1, 5, 6, 4}

	ans := findKthLargestII(nums, 2)

	assert.Check(t, ans == 5)
}
