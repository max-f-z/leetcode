package main

func minTime(n int, edges [][]int, hasApple []bool) int {
	graph := map[int][]int{}

	for _, v := range edges {
		if graph[v[0]] == nil {
			graph[v[0]] = []int{}
		}
		if graph[v[1]] == nil {
			graph[v[1]] = []int{}
		}
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	visited := make([]int, n)

	return minTimeHelper(0, graph, hasApple, &visited)
}

func minTimeHelper(cur int, graph map[int][]int, apples []bool, visited *[]int) int {
	sum := 0

	(*visited)[cur] = 1

	for _, v := range graph[cur] {
		if (*visited)[v] == 1 {
			continue
		}
		sum += minTimeHelper(v, graph, apples, visited)
	}

	if cur == 0 {
		return sum
	}

	if sum == 0 && !apples[cur] {
		return 0
	}

	return sum + 2
}
