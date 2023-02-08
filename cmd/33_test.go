package main

import "testing"

func Test33(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	ans := search33(arr, target)
	t.Log(ans)
}
