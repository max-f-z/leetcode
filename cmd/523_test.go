package main

import "testing"

func Test523(t *testing.T) {
	nums := []int{23, 2, 4, 6, 6}
	k := 7

	ans := checkSubarraySum(nums, k)

	t.Log(ans)
}
