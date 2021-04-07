package main

import "math"

func minDistance573(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	distTree := make([]int, len(nuts))
	distSquirrel := make([]int, len(nuts))

	sum := 0
	res := math.MaxInt32

	for i := 0; i < len(nuts); i++ {
		distTree[i] = int(math.Abs(float64(nuts[i][0]-tree[0]))) + int(math.Abs(float64(nuts[i][1]-tree[1])))

		sum += distTree[i]
	}

	for i := 0; i < len(nuts); i++ {
		distSquirrel[i] = int(math.Abs(float64(nuts[i][0]-squirrel[0]))) + int(math.Abs(float64(nuts[i][1]-squirrel[1])))

		tmp := distSquirrel[i] + sum*2 - distTree[i]

		if res > tmp {
			res = tmp
		}
	}

	return res
}
