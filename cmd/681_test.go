package main

import "testing"

func Test681(t *testing.T) {
	time := nextClosestTime("23:59")
	t.Log(time)
}
