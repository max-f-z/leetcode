package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2418(t *testing.T) {
	names := []string{"Alice", "Bob", "Bob"}
	heights := []int{155, 185, 150}

	ans := sortPeople(names, heights)
	assert.DeepEqual(t, ans, []string{"Bob", "Alice", "Bob"})
}
