package main

import (
	"container/heap"
	"strings"
)

type charCnt struct {
	char byte
	cnt  int
}

type maxHeap767 []*charCnt

func (mh maxHeap767) Len() int {
	return len(mh)
}

func (mh maxHeap767) Swap(i, j int) {
	(mh)[i], (mh)[j] = (mh)[j], (mh)[i]
}

func (mh maxHeap767) Less(i, j int) bool {
	return mh[i].cnt > mh[j].cnt
}

func (mh *maxHeap767) Push(x interface{}) {
	v := x.(*charCnt)
	*mh = append(*mh, v)
}

func (mh *maxHeap767) Pop() (v interface{}) {
	v, *mh = (*mh)[len(*mh)-1], (*mh)[:len(*mh)-1]
	return v
}

func (mh maxHeap767) Top() *charCnt {
	return (mh)[0]
}

func reorganizeString(S string) string {
	cols := map[byte]int{}

	for i := 0; i < len(S); i++ {
		cols[S[i]]++
	}

	arr := make([]*charCnt, len(cols))

	idx := 0

	for k, v := range cols {
		cc := &charCnt{
			char: k,
			cnt:  v,
		}
		arr[idx] = cc
		idx++
	}

	mh := maxHeap767(arr)
	heap.Init(&mh)

	sb := strings.Builder{}

	v := heap.Pop(&mh)
	cc := v.(*charCnt)

	prev := cc.char
	sb.WriteByte(cc.char)
	cc.cnt--

	if cc.cnt > 0 {
		heap.Push(&mh, cc)
	}

	for mh.Len() > 0 {
		v := heap.Pop(&mh)
		cc := v.(*charCnt)

		if prev != cc.char {
			sb.WriteByte(cc.char)
			cc.cnt--
			prev = cc.char

			if cc.cnt > 0 {
				heap.Push(&mh, cc)
			}
			continue
		} else {
			if mh.Len() == 0 {
				return ""
			}

			vv := heap.Pop(&mh)
			ccc := vv.(*charCnt)

			sb.WriteByte(ccc.char)
			ccc.cnt--
			if ccc.cnt > 0 {
				heap.Push(&mh, ccc)
			}

			sb.WriteByte(cc.char)
			cc.cnt--
			prev = cc.char

			if cc.cnt > 0 {
				heap.Push(&mh, cc)
			}
		}
	}

	return sb.String()
}
