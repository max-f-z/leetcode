package main

import (
	"container/heap"
	"strings"
)

func longestDiverseString(a int, b int, c int) string {
	mh := maxHeap1405{}
	if a > 0 {
		heap.Push(&mh, &mhEle{
			char:  'a',
			count: a,
		})
	}

	if b > 0 {
		heap.Push(&mh, &mhEle{
			char:  'b',
			count: b,
		})
	}

	if c > 0 {
		heap.Push(&mh, &mhEle{
			char:  'c',
			count: c,
		})
	}

	heap.Init(&mh)
	ans := strings.Builder{}

	var prefix byte

	for mh.Len() > 0 {
		first := heap.Pop(&mh).(*mhEle)

		if mh.Len() == 0 {
			if prefix == first.char {
				break
			}

			if first.count >= 2 {
				ans.WriteByte(first.char)
				ans.WriteByte(first.char)
				first.count -= 2
				prefix = first.char
			} else if first.count == 1 {
				ans.WriteByte(first.char)
				first.count--
				prefix = first.char
			}
		} else {
			if prefix == first.char {
				second := heap.Pop(&mh).(*mhEle)
				ans.WriteByte(second.char)
				second.count--
				prefix = second.char

				if second.count > 0 {
					heap.Push(&mh, second)
				}
			} else {
				if first.count >= 2 {
					ans.WriteByte(first.char)
					ans.WriteByte(first.char)
					first.count -= 2
					prefix = first.char
				} else if first.count == 1 {
					ans.WriteByte(first.char)
					first.count--
					prefix = first.char
				}
			}
		}

		if first.count > 0 {
			heap.Push(&mh, first)
		}
	}

	return ans.String()
}

type maxHeap1405 []*mhEle

type mhEle struct {
	char  byte
	count int
}

func (mh maxHeap1405) Less(i, j int) bool {
	return mh[i].count > mh[j].count
}

func (mh maxHeap1405) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh maxHeap1405) Len() int {
	return len(mh)
}

func (mh *maxHeap1405) Push(x interface{}) {
	*mh = append(*mh, x.(*mhEle))
}

func (mh *maxHeap1405) Pop() (v interface{}) {
	(*mh), v = (*mh)[:mh.Len()-1], (*mh)[mh.Len()-1]
	return v
}

func (mh maxHeap1405) Top() *mhEle {
	return mh[0]
}
