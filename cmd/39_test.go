package main

import "testing"

func Test39(t *testing.T) {
	candidates := []int{2, 3, 6, 7}

	ans := combinationSumII(candidates, 7)

	t.Log(ans)
}
