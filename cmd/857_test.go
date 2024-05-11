package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test857(t *testing.T) {
	quality := []int{10, 20, 5}
	wage := []int{70, 50, 30}
	ans := mincostToHireWorkers(quality, wage, 2)
	assert.Check(t, ans == 105.0)
}
