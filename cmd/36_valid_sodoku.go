package main

import "fmt"

// https://www.youtube.com/watch?v=Pl7mMcBm2b8

func isValidSudoku(board [][]byte) bool {

	box := map[string]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if !box[fmt.Sprintf("Row:%d-%c", i, board[i][j])] && !box[fmt.Sprintf("Col:%d-%c", j, board[i][j])] && !box[fmt.Sprintf("Box:%d-%d-%c", i/3, j/3, board[i][j])] {
					box[fmt.Sprintf("Row:%d-%c", i, board[i][j])] = true
					box[fmt.Sprintf("Col:%d-%c", j, board[i][j])] = true
					box[fmt.Sprintf("Box:%d-%d-%c", i/3, j/3, board[i][j])] = true
					continue
				}
				return false
			}
		}
	}

	return true
}
