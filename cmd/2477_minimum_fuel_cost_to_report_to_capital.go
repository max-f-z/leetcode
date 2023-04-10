package main

//lint:ignore U1000 unused
func minimumFuelCost(roads [][]int, seats int) int64 {
	if len(roads) == 0 {
		return 0
	}

	graph := map[int][]int{}
	for _, v := range roads {
		if graph[v[0]] == nil {
			graph[v[0]] = []int{}
		}
		if graph[v[1]] == nil {
			graph[v[1]] = []int{}
		}
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	visited := make([]int, len(graph))
	visited[0] = 1
	ans, _ := minimumFuelCostDFS(graph, 0, &visited, seats)

	return int64(ans)
}

func minimumFuelCostDFS(graph map[int][]int, current int, visited *[]int, seats int) (int, int) {
	reps := 1
	fuel := 0

	(*visited)[current] = 1

	if current != 0 && len(graph[current]) == 1 {
		return 1, 1
	}

	for _, v := range graph[current] {
		if (*visited)[v] == 0 {
			(*visited)[v] = 1
			childFuel, childReps := minimumFuelCostDFS(graph, v, visited, seats)
			reps += childReps
			fuel += childFuel
		}
	}

	if current != 0 {
		if reps%seats == 0 {
			fuel += reps / seats
		} else {
			fuel += reps/seats + 1
		}
	}

	return fuel, reps
}
