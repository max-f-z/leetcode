package main

import (
	"container/heap"
	"sort"
	"strings"
)

func addBoldTag(s string, dict []string) string {
	sb := strings.Builder{}

	segs := [][]int{}

	for _, v := range dict {
		idx := findIndices(s, v)
		segs = append(segs, idx...)
	}

	sort.Slice(segs, func(i, j int) bool {
		if segs[i][0] != segs[j][0] {
			return segs[i][0] < segs[j][0]
		}
		return segs[i][1] < segs[j][1]
	})

	merged := [][]int{}

	mh := &minHeap{}
	heap.Init(mh)

	start := -1

	for idx, v := range segs {
		if idx == 0 {
			start = v[0]
			heap.Push(mh, v[1])
			continue
		}

		if mh.Top()+1 < v[0] {
			merged = append(merged, []int{start, mh.Top()})
			heap.Pop(mh)
			heap.Push(mh, v[1])
			start = v[0]
			continue
		}

		if mh.Top() < v[1] {
			heap.Pop(mh)
			heap.Push(mh, v[1])
			continue
		}
	}

	if mh.Len() > 0 {
		merged = append(merged, []int{start, mh.Top()})
	}

	start = 0

	for _, v := range merged {
		sb.WriteString(s[start:v[0]])
		sb.WriteString("<b>")
		sb.WriteString(s[v[0] : v[1]+1])
		sb.WriteString("</b>")
		start = v[1] + 1
	}

	if start < len(s) {
		sb.WriteString(s[start:])
	}

	return sb.String()
}

func findIndices(s string, dict string) [][]int {
	res := [][]int{}

	for i := 0; i <= len(s)-len(dict); i++ {
		if s[i:i+len(dict)] == dict {
			res = append(res, []int{i, i + len(dict) - 1})
		}
	}

	return res
}
