package main

// https://www.youtube.com/watch?v=ZV0InYdJKYw
//
//lint:ignore U1000 unused
func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := '1'; k <= '9'; k++ {
					if valid(board, i, j, byte(k)) {
						board[i][j] = byte(k)
						if solve(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func valid(board [][]byte, row, col int, k byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}

		if board[i][col] == k {
			return false
		}

		if board[3*(row/3)+i/3][3*(col/3)+i%3] == k {
			return false
		}
	}

	return true
}
