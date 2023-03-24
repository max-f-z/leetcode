package main

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	numOfGroups := 0
	visited := map[int]bool{}

	graph := map[int][]int{}

	for i := 0; i < len(connections); i++ {
		if graph[connections[i][0]] == nil {
			graph[connections[i][0]] = []int{}
		}
		if graph[connections[i][1]] == nil {
			graph[connections[i][1]] = []int{}
		}
		graph[connections[i][0]] = append(graph[connections[i][0]], connections[i][1])
		graph[connections[i][1]] = append(graph[connections[i][1]], connections[i][0])
	}

	h := &helper1319{
		graph:   graph,
		visited: visited,
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			h.visited[i] = true
			h.dfs(i)
			numOfGroups++
		}
	}

	return numOfGroups - 1
}

type helper1319 struct {
	graph   map[int][]int
	visited map[int]bool
}

func (h *helper1319) dfs(cur int) {
	for i := 0; i < len(h.graph[cur]); i++ {
		if !h.visited[h.graph[cur][i]] {
			h.visited[h.graph[cur][i]] = true
			h.dfs(h.graph[cur][i])
		}
	}
}
