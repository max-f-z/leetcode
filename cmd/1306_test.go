package main

import "testing"

func Test1306(t *testing.T) {
	board := []int{4, 2, 3, 0, 3, 1, 2}

	res := canReach(board, 5)

	t.Log(res)
}
