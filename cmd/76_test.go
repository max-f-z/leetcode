package main

import "testing"

func Test76(t *testing.T) {
	a := "ADOBECODEBANC"
	b := "ABC"
	c := minWindow(a, b)

	t.Log(c)
}
