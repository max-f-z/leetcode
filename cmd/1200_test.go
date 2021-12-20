package main

import "testing"

func Test1200(t *testing.T) {
	arr := []int{1, 3, 6, 10, 15}

	ans := minimumAbsDifference(arr)

	t.Log(ans)
}
