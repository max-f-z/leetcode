package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test85(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	assert.Equal(t, maximalRectangle(matrix), 6)
}
