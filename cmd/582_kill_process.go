package main

import (
	"container/list"
)

//lint:ignore U1000 unused
func killProcess(pid []int, ppid []int, kill int) []int {
	cols := map[int][]int{}

	for i := 0; i < len(pid); i++ {
		cols[pid[i]] = []int{}
	}

	for i := 0; i < len(ppid); i++ {
		cols[ppid[i]] = append(cols[ppid[i]], pid[i])
	}

	q := list.New()

	res := []int{}

	q.PushBack(kill)
	for q.Len() > 0 {
		val := q.Front()
		id := val.Value.(int)
		res = append(res, id)
		q.Remove(val)

		for _, v := range cols[id] {
			q.PushBack(v)
		}
	}

	return res
}
