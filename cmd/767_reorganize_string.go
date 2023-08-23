package main

import (
	"container/heap"
	"strings"
)

type charCnt struct {
	char byte
	cnt  int
}

type mhCharCnt []*charCnt

func (mh mhCharCnt) Len() int {
	return len(mh)
}

func (mh mhCharCnt) Less(i, j int) bool {
	return mh[i].cnt > mh[j].cnt
}

func (mh mhCharCnt) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *mhCharCnt) Push(x interface{}) {
	*mh = append(*mh, x.(*charCnt))
}

func (mh *mhCharCnt) Pop() (v interface{}) {
	v, *mh = (*mh)[mh.Len()-1], (*mh)[:mh.Len()-1]
	return v
}

func reorganizeString(S string) string {
	cnt := map[byte]int{}

	for i := 0; i < len(S); i++ {
		cnt[S[i]]++
	}

	num := len(cnt)
	idx := 0
	chars := make([]*charCnt, num)
	for k, v := range cnt {
		cc := &charCnt{
			char: k,
			cnt:  v,
		}
		chars[idx] = cc
		idx++
	}

	mhcc := mhCharCnt(chars)
	heap.Init(&mhcc)

	sb := strings.Builder{}

	cc := heap.Pop(&mhcc).(*charCnt)
	prev := cc.char
	sb.WriteByte(cc.char)
	cc.cnt--

	if cc.cnt > 0 {
		heap.Push(&mhcc, cc)
	}

	for mhcc.Len() > 0 {
		cc := heap.Pop(&mhcc).(*charCnt)
		if prev == cc.char {
			if mhcc.Len() == 0 {
				return ""
			}

			ccc := heap.Pop(&mhcc).(*charCnt)
			sb.WriteByte(ccc.char)
			ccc.cnt--

			if ccc.cnt > 0 {
				heap.Push(&mhcc, ccc)
			}
		}
		sb.WriteByte(cc.char)
		cc.cnt--
		prev = cc.char
		if cc.cnt > 0 {
			heap.Push(&mhcc, cc)
		}
	}

	return sb.String()
}
