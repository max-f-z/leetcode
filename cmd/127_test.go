package main

import "testing"

func Test127(t *testing.T) {
	list := []string{"cog"}
	val := ladderLength("hog", "cog", list)

	t.Log(val)
}
