package main

func shortestDistance505(maze [][]int, start []int, destination []int) int {
	h := len(maze)
	w := len(maze[0])

	// x, y, previous direction, direction, distance
	start = append(start, 0, 0, 0)
	maze[start[0]][start[1]] = -1
	queue := [][]int{start}

	for len(queue) > 0 {
		newQueue := [][]int{}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if v[0] == destination[0] && v[1] == destination[1] && v[3] == 0 {
				return v[4]
			}

			// left
			if (v[3] == 1 || v[3] == 0) && v[2] != 1 {
				if v[1]-1 >= 0 && maze[v[0]][v[1]-1] != 1 {
					newQueue = append(newQueue, []int{v[0], v[1] - 1, 2, 1, v[4] + 1})
				} else {
					if v[3] != 0 && maze[v[0]][v[1]] != -1 {
						queue = append(queue, []int{v[0], v[1], 2, 0, v[4]})
						maze[v[0]][v[1]] = -1
					}
				}
			}

			// right
			if (v[3] == 2 || v[3] == 0) && v[2] != 2 {
				if v[1]+1 < w && maze[v[0]][v[1]+1] != 1 {
					newQueue = append(newQueue, []int{v[0], v[1] + 1, 1, 2, v[4] + 1})
				} else {
					if v[3] != 0 && maze[v[0]][v[1]] != -1 {
						queue = append(queue, []int{v[0], v[1], 1, 0, v[4]})
						maze[v[0]][v[1]] = -1
					}
				}
			}

			// up
			if (v[3] == 3 || v[3] == 0) && v[2] != 3 {
				if v[0]-1 >= 0 && maze[v[0]-1][v[1]] != 1 {
					newQueue = append(newQueue, []int{v[0] - 1, v[1], 4, 3, v[4] + 1})
				} else {
					if v[3] != 0 && maze[v[0]][v[1]] != -1 {
						queue = append(queue, []int{v[0], v[1], 4, 0, v[4]})
						maze[v[0]][v[1]] = -1
					}
				}
			}

			// down
			if (v[3] == 4 || v[3] == 0) && v[2] != 4 {
				if v[0]+1 < h && maze[v[0]+1][v[1]] != 1 {
					newQueue = append(newQueue, []int{v[0] + 1, v[1], 3, 4, v[4] + 1})
				} else {
					if v[3] != 0 && maze[v[0]][v[1]] != -1 {
						queue = append(queue, []int{v[0], v[1], 3, 0, v[4]})
						maze[v[0]][v[1]] = -1
					}
				}
			}

		}

		queue = newQueue
	}

	return -1
}
