package main

import "math"

func minArea(image [][]byte, x int, y int) int {
	minX, minY := math.MaxInt32, math.MaxInt32
	maxX, maxY := -1, -1

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j] == '1' {
				if i > maxX {
					maxX = i
				}
				if i < minX {
					minX = i
				}

				if j > maxY {
					maxY = j
				}
				if j < minY {
					minY = j
				}
			}
		}
	}

	return (maxX - minX + 1) * (maxY - minY + 1)
}
