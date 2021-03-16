package main

import (
	"container/heap"
	"strings"
)

type maxHeapEl struct {
	char  byte
	count int
}

type maxHeap358 []maxHeapEl

func (h maxHeap358) Len() int {
	return len(h)
}

func (h maxHeap358) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count > h[j].count
	}
	return h[i].char < h[j].char
}

func (h maxHeap358) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap358) Push(x interface{}) {
	v, _ := x.(maxHeapEl)
	*h = append(*h, v)
}

func (h *maxHeap358) Pop() (v interface{}) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return v
}

func rearrangeString(s string, k int) string {
	chars := map[byte]int{}

	if k == 0 {
		return s
	}

	for i := range s {
		chars[s[i]]++
	}

	mh := &maxHeap358{}
	heap.Init(mh)

	for k, v := range chars {
		heap.Push(mh, maxHeapEl{char: k, count: v})
	}

	res := strings.Builder{}

	for mh.Len() > 0 {
		if len(s)-res.Len() > mh.Len() && mh.Len() < k {
			return ""
		}
		pushBack := []maxHeapEl{}

		for i := k; i > 0; i-- {
			if mh.Len() == 0 {
				break
			}
			tmp := heap.Pop(mh).(maxHeapEl)
			tmp.count--
			if tmp.count > 0 {
				pushBack = append(pushBack, tmp)
			}
			res.WriteByte(tmp.char)
		}

		for i := 0; i < len(pushBack); i++ {
			heap.Push(mh, pushBack[i])
		}
	}

	return res.String()
}
