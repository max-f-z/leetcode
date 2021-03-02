package main

import (
	"strconv"
	"strings"
)

func slidingPuzzle(board [][]int) int {
	oringinal := [][]int{{1, 2, 3}, {4, 5, 0, 0}}

	oringinalStr := boardToStr(oringinal)
	visited := map[string]bool{}
	visited[oringinalStr] = true
	targetStr := boardToStr(board)

	if oringinalStr == targetStr {
		return 0
	}

	queue := [][][]int{oringinal}

	for len(queue) > 0 {
		newQ := [][][]int{}
		for _, v := range queue {
			var i, j int
			found := false
			for i = 0; i < 2; i++ {
				for j = 0; j < 3; j++ {
					if v[i][j] == 0 {
						found = true
						break
					}
				}
				if found {
					break
				}
			}

			// down
			if i+1 < 2 {
				tmp := dupBoard(v)
				tmp[i][j] = tmp[i+1][j]
				tmp[i+1][j] = 0

				tmpStr := boardToStr(tmp)
				if tmpStr == targetStr {
					return v[1][3] + 1
				}

				if !visited[tmpStr] {
					visited[tmpStr] = true
					tmp[1] = append(tmp[1], v[1][3]+1)
					newQ = append(newQ, tmp)
				}
			}

			// up
			if i-1 >= 0 {
				tmp := dupBoard(v)
				tmp[i][j] = tmp[i-1][j]
				tmp[i-1][j] = 0
				tmpStr := boardToStr(tmp)
				if tmpStr == targetStr {
					return v[1][3] + 1
				}

				if !visited[tmpStr] {
					visited[tmpStr] = true
					tmp[1] = append(tmp[1], v[1][3]+1)
					newQ = append(newQ, tmp)
				}
			}

			// left
			if j+1 < 3 {
				tmp := dupBoard(v)
				tmp[i][j] = tmp[i][j+1]
				tmp[i][j+1] = 0
				tmpStr := boardToStr(tmp)
				if tmpStr == targetStr {
					return v[1][3] + 1
				}

				if !visited[tmpStr] {
					visited[tmpStr] = true
					tmp[1] = append(tmp[1], v[1][3]+1)
					newQ = append(newQ, tmp)
				}
			}

			// right
			if j-1 >= 0 {
				tmp := dupBoard(v)
				tmp[i][j] = tmp[i][j-1]
				tmp[i][j-1] = 0
				tmpStr := boardToStr(tmp)
				if tmpStr == targetStr {
					return v[1][3] + 1
				}

				if !visited[tmpStr] {
					visited[tmpStr] = true
					tmp[1] = append(tmp[1], v[1][3]+1)
					newQ = append(newQ, tmp)
				}
			}

		}

		queue = newQ
	}

	return -1
}

func boardToStr(board [][]int) string {
	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(board[0][0]))
	builder.WriteString(strconv.Itoa(board[0][1]))
	builder.WriteString(strconv.Itoa(board[0][2]))
	builder.WriteString(strconv.Itoa(board[1][0]))
	builder.WriteString(strconv.Itoa(board[1][1]))
	builder.WriteString(strconv.Itoa(board[1][2]))
	return builder.String()
}

func dupBoard(board [][]int) [][]int {
	dup := make([][]int, 2)

	for i := 0; i < 2; i++ {
		dup[i] = []int{}
		for j := 0; j < 3; j++ {
			dup[i] = append(dup[i], board[i][j])
		}
	}

	return dup
}
