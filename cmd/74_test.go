package main

import "testing"

func Test74(t *testing.T) {
	// matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	matrix := [][]int{{1}, {3}}
	target := 3

	ans := searchMatrix(matrix, target)

	t.Log(ans)
}
