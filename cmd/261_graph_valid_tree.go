package main

// https://www.youtube.com/watch?v=n_t0a_8H8VY

func validTree(n int, edges [][]int) bool {
	graph := make(map[int][]int)

	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	vis := make([]bool, n)

	if !validTreeHelper(vis, graph, 0, -1) {
		return false
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			return false
		}
	}

	return true
}

func validTreeHelper(vis []bool, graph map[int][]int, start int, parent int) bool {
	if vis[start] {
		return false
	}

	vis[start] = true

	for _, v := range graph[start] {
		if v != parent {
			if !validTreeHelper(vis, graph, v, start) {
				return false
			}
		}
	}

	return true
}
