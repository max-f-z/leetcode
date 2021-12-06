package main

import "testing"

func Test239(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	ans := maxSlidingWindow(nums, k)

	t.Log(ans)
}
