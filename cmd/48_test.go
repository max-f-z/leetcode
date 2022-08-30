package main

import (
	"fmt"
	"testing"
)

func Test48(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	rotate(matrix)

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	t.Log(matrix)
}
