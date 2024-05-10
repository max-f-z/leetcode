package main

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	indexedQueries := make([][3]int, len(queries))

	for i := 0; i < len(queries); i++ {
		indexedQueries[i] = [3]int{i, queries[i] - x, queries[i]}
	}

	sort.Slice(indexedQueries, func(i, j int) bool {
		return indexedQueries[i][1] < indexedQueries[j][1]
	})

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	ans := make([]int, len(queries))

	l, r := 0, 0
	active := map[int]int{}
	for i := 0; i < len(indexedQueries); i++ {
		idx, start, end := indexedQueries[i][0], indexedQueries[i][1], indexedQueries[i][2]

		for r < len(logs) && logs[r][1] <= end {
			active[logs[r][0]]++
			r++
		}

		for l < len(logs) && logs[l][1] < start {
			active[logs[l][0]]--
			if active[logs[l][0]] == 0 {
				delete(active, logs[l][0])
			}
			l++
		}

		ans[idx] = n - len(active)
	}

	return ans
}
