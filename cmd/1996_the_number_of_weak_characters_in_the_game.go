package main

import (
	"sort"
)

type characters [][]int

func (c characters) Len() int {
	return len(c)
}

func (c characters) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c characters) Less(i, j int) bool {
	if c[i][0] == c[j][0] {
		return c[i][1] > c[j][1]
	}
	return c[i][0] < c[j][0]
}

func numberOfWeakCharacters(properties [][]int) int {
	c := characters(properties)
	sort.Sort(c)

	weak := 0
	maxDefense := 0

	for i := len(c) - 1; i >= 0; i-- {
		if c[i][1] < maxDefense {
			weak++
		} else {
			maxDefense = c[i][1]
		}
	}

	return weak
}
