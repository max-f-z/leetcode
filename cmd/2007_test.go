package main

import "testing"

func Test2007(t *testing.T) {
	changed := []int{1, 3, 4, 2, 6, 8}

	ans := findOriginalArray(changed)

	t.Log(ans)
}
