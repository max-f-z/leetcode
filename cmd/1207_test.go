package main

import "testing"

func Test1207(t *testing.T) {
	arr := []int{1, 2, 2, 1, 1, 3}

	ans := uniqueOccurrences(arr)

	t.Log(ans)
}
