package main

import "testing"

func Test547(t *testing.T) {

	val := findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}})

	t.Log(val)
}
