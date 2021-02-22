package main

func minKnightMoves(x int, y int) int {

	if x == 0 && y == 0 {
		return 0
	}

	board := make(map[int]map[int]bool)
	board[0] = map[int]bool{}
	board[0][0] = true

	queue := [][]int{{0, 0, 0}}

	for {
		newQ := [][]int{}

		for i := 0; i < len(queue); i++ {
			if addToBoard(board, queue[i][0]+1, queue[i][1]+2) {
				newQ = append(newQ, []int{queue[i][0] + 1, queue[i][1] + 2, queue[i][2] + 1})
				if queue[i][0]+1 == x && queue[i][1]+2 == y {
					return queue[i][2] + 1
				}
			}

			if addToBoard(board, queue[i][0]+2, queue[i][1]+1) {
				newQ = append(newQ, []int{queue[i][0] + 2, queue[i][1] + 1, queue[i][2] + 1})
				if queue[i][0]+2 == x && queue[i][1]+1 == y {
					return queue[i][2] + 1
				}
			}

			if addToBoard(board, queue[i][0]+1, queue[i][1]-2) {
				newQ = append(newQ, []int{queue[i][0] + 1, queue[i][1] - 2, queue[i][2] + 1})
				if queue[i][0]+1 == x && queue[i][1]-2 == y {
					return queue[i][2] + 1
				}
			}

			if addToBoard(board, queue[i][0]+2, queue[i][1]-1) {
				newQ = append(newQ, []int{queue[i][0] + 2, queue[i][1] - 1, queue[i][2] + 1})
				if queue[i][0]+2 == x && queue[i][1]-1 == y {
					return queue[i][2] + 1
				}
			}

			if addToBoard(board, queue[i][0]-1, queue[i][1]+2) {
				newQ = append(newQ, []int{queue[i][0] - 1, queue[i][1] + 2, queue[i][2] + 1})
				if queue[i][0]-1 == x && queue[i][1]+2 == y {
					return queue[i][2] + 1
				}
			}

			if addToBoard(board, queue[i][0]-1, queue[i][1]-2) {
				newQ = append(newQ, []int{queue[i][0] - 1, queue[i][1] - 2, queue[i][2] + 1})
				if queue[i][0]-1 == x && queue[i][1]-2 == y {
					return queue[i][2] + 1
				}
			}

			if addToBoard(board, queue[i][0]-2, queue[i][1]+1) {
				newQ = append(newQ, []int{queue[i][0] - 2, queue[i][1] + 1, queue[i][2] + 1})
				if queue[i][0]-2 == x && queue[i][1]+1 == y {
					return queue[i][2] + 1
				}
			}

			if addToBoard(board, queue[i][0]-2, queue[i][1]-1) {
				newQ = append(newQ, []int{queue[i][0] - 2, queue[i][1] - 1, queue[i][2] + 1})
				if queue[i][0]-2 == x && queue[i][1]-1 == y {
					return queue[i][2] + 1
				}
			}
		}

		queue = newQ
	}

}

func addToBoard(board map[int]map[int]bool, x, y int) bool {
	if board[x] == nil {
		board[x] = make(map[int]bool)
	}
	if board[x][y] {
		return false
	}
	board[x][y] = true
	return true
}
