package main

//lint:ignore U1000 unused
func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}

	h := len(rooms)
	w := len(rooms[0])

	gates := [][]int{}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if rooms[i][j] == 0 {
				gates = append(gates, []int{i, j})
			}
		}
	}

	for i := 0; i < len(gates); i++ {
		visited := make([][]int, h)

		for i := 0; i < h; i++ {
			visited[i] = make([]int, w)
		}
		stack := [][]int{{gates[i][0], gates[i][1]}}
		visited[gates[i][0]][gates[i][1]] = 1
		for len(stack) != 0 {
			visit(rooms, w, h, &stack, visited)
		}
	}
}

func visit(rooms [][]int, w, h int, stack *[][]int, visited [][]int) {
	if len(*stack) == 0 {
		return
	}

	x, y := (*stack)[0][0], (*stack)[0][1]
	*stack = (*stack)[1:]

	if y+1 < w {
		if rooms[x][y+1] != -1 && rooms[x][y]+1 < rooms[x][y+1] && visited[x][y+1] == 0 {
			rooms[x][y+1] = rooms[x][y] + 1
			(*stack) = append((*stack), []int{x, y + 1})
			visited[x][y+1] = 1
		}
	}

	if x+1 < h {
		if rooms[x+1][y] != -1 && rooms[x][y]+1 < rooms[x+1][y] && visited[x+1][y] == 0 {
			rooms[x+1][y] = rooms[x][y] + 1
			(*stack) = append((*stack), []int{x + 1, y})
			visited[x+1][y] = 1
		}
	}

	if y-1 > -1 {
		if rooms[x][y-1] != -1 && rooms[x][y]+1 < rooms[x][y-1] && visited[x][y-1] == 0 {
			rooms[x][y-1] = rooms[x][y] + 1
			(*stack) = append((*stack), []int{x, y - 1})
			visited[x][y-1] = 1
		}
	}

	if x-1 > -1 {
		if rooms[x-1][y] != -1 && rooms[x][y]+1 < rooms[x-1][y] && visited[x-1][y] == 0 {
			rooms[x-1][y] = rooms[x][y] + 1
			(*stack) = append((*stack), []int{x - 1, y})
			visited[x-1][y] = 1
		}
	}
}
