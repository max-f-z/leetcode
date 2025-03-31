package main

import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)

	adj := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		adj[i] = weights[i] + weights[i+1]
	}

	sort.Ints(adj)

	ans := int64(0)

	for i := 0; i < k-1; i++ {
		ans += int64(adj[n-2-i] - adj[i])
	}

	return ans
}
