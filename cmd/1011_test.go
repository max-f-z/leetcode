package main

import "testing"

func Test1011(t *testing.T) {
	w := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ans := shipWithinDays(w, 5)
	t.Log(ans)
}
