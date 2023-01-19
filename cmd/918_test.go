package main

import "testing"

func Test918(t *testing.T) {
	nums := []int{1, -2, 3, -2}

	ans := maxSubarraySumCircular(nums)

	t.Log(ans)
}
