package main

import (
	"math"
)

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	hist := make([][]int, len(matrix))
	for i := range hist {
		hist[i] = make([]int, len(matrix[0]))
	}

	area := 0

	for i := range matrix {
		for k := range matrix[i] {
			if matrix[i][k] == '1' {
				if i == 0 {
					hist[i][k]++
				} else {
					hist[i][k] = hist[i-1][k] + 1
				}
			} else {
				hist[i][k] = 0
			}
		}
		temp := histogram(hist[i])
		if area < temp {
			area = temp
		}
	}
	return area
}

func histogram(m []int) int {
	max := m[0]
	area := max

	for i := 1; i < len(m); i++ {
		if i+1 < len(m) && m[i] <= m[i+1] {
			continue
		}
		min := math.MaxInt32
		for j := i; j >= 0; j-- {
			if m[j] < min {
				min = m[j]
			}
			temp := min * (i - j + 1)
			if temp > area {
				area = temp
			}
		}

	}

	return area
}
