package main

import "sort"

type height struct {
	Name   string
	Height int
}

func sortPeople(names []string, heights []int) []string {
	hs := make([]*height, len(names))

	for i := 0; i < len(names); i++ {
		hs[i] = &height{
			Name:   names[i],
			Height: heights[i],
		}
	}

	sort.Slice(hs, func(i, j int) bool {
		return hs[i].Height > hs[j].Height
	})

	ans := make([]string, len(names))

	for i := 0; i < len(names); i++ {
		ans[i] = hs[i].Name
	}

	return ans
}
