package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2747(t *testing.T) {
	logs := [][]int{{1, 3}, {2, 6}, {1, 5}}
	queries := []int{10, 11}

	ans := countServers(3, logs, 5, queries)
	assert.DeepEqual(t, ans, []int{1, 2})
}
