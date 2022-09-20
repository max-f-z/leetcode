package main

import "testing"

func Test718(t *testing.T) {
	a := []int{1, 2, 3, 2, 8}
	b := []int{5, 6, 1, 4, 7}

	res := findLength(a, b)

	t.Log(res)
}
