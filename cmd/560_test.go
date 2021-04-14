package main

import "testing"

func Test560(t *testing.T) {
	nums := []int{1, 2, 1, 2, 1}

	n := subarraySum(nums, 3)

	t.Log(n)
}
