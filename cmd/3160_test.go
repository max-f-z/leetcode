package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3160(t *testing.T) {
	limit := 4
	queries := [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}

	ans := queryResults(limit, queries)

	assert.DeepEqual(t, ans, []int{1, 2, 2, 3})

	limit = 4
	queries = [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}

	ans = queryResults(limit, queries)

	assert.DeepEqual(t, ans, []int{1, 2, 2, 3, 4})
}
