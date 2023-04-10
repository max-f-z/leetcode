package main

func hasPath(maze [][]int, start []int, destination []int) bool {
	var directions = [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	m, n := len(maze), len(maze[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	visited[start[0]][start[1]] = true
	queue := [][]int{start}
	for len(queue) > 0 {
		newQueue := [][]int{}
		for _, position := range queue {
			if position[0] == destination[0] && position[1] == destination[1] {
				return true
			}

			for _, direction := range directions {
				xx, yy := position[0], position[1]
				for xx+direction[0] >= 0 && xx+direction[0] < m &&
					yy+direction[1] >= 0 && yy+direction[1] < n &&
					maze[xx+direction[0]][yy+direction[1]] == 0 {
					xx, yy = xx+direction[0], yy+direction[1]
				}

				if visited[xx][yy] {
					continue
				}
				visited[xx][yy] = true
				newQueue = append(newQueue, []int{xx, yy})
			}
		}
		queue = newQueue
	}

	return false
}
