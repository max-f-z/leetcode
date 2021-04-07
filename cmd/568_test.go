package main

import "testing"

func Test568(t *testing.T) {
	flights := [][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}
	days := [][]int{{1, 3, 1}, {6, 0, 3}, {3, 3, 3}}
	res := maxVacationDays(flights, days)

	t.Log(res)
}
