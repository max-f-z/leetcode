package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test983(t *testing.T) {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}

	ans := mincostTickets(days, costs)

	assert.Check(t, ans == 11)
}
