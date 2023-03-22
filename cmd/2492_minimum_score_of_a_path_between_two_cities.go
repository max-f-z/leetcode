package main

func minScore(n int, roads [][]int) int {
	graph := map[int]map[int]int{}

	for _, road := range roads {
		if graph[road[0]] == nil {
			graph[road[0]] = make(map[int]int)
		}
		if graph[road[1]] == nil {
			graph[road[1]] = make(map[int]int)
		}

		graph[road[0]][road[1]] = road[2]
		graph[road[1]][road[0]] = road[2]
	}

	min := 10001

	queue := []int{1}
	visited := map[int]bool{}
	visited[1] = true

	for len(queue) > 0 {
		newQ := []int{}
		for _, node := range queue {
			for k, v := range graph[node] {
				if !visited[k] {
					newQ = append(newQ, k)
					visited[k] = true
				}
				if v < min {
					min = v
				}
			}
		}
		queue = newQ
	}

	return min
}
