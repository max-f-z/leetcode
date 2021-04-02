package main

import "testing"

func TestSplitLoopedString(t *testing.T) {
	strs := []string{"lc", "evol", "cdy"}
	res := splitLoopedString(strs)

	t.Log(res)
}
