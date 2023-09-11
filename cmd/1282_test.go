package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1282(t *testing.T) {
	groupSize := []int{3, 3, 3, 3, 3, 1, 3}

	ans := groupThePeople(groupSize)

	assert.Check(t, len(ans) == 3)
}
