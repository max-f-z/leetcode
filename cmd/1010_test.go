package main

import "testing"

func Test1010(t *testing.T) {
	ways := numPairsDivisibleBy60([]int{60, 60, 60})

	t.Log(ways)
}
