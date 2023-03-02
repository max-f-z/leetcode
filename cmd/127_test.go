package main

import "testing"

func Test127(t *testing.T) {
	list := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	val := ladderLengthPlus("hit", "cog", list)

	t.Log(val)
}
