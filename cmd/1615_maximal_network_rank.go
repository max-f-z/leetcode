package main

func maximalNetworkRank(n int, roads [][]int) int {
	graph := map[int]map[int]bool{}

	for i := 0; i < len(roads); i++ {
		a, b := roads[i][0], roads[i][1]

		if graph[a] == nil {
			graph[a] = make(map[int]bool)
		}

		if graph[b] == nil {
			graph[b] = make(map[int]bool)
		}

		graph[a][b] = true
		graph[b][a] = true
	}
	max := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			val := len(graph[i]) + len(graph[j])
			if graph[i][j] {
				val--
			}

			if val > max {
				max = val
			}
		}
	}

	return max
}
