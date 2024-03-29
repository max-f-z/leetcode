package main

import (
	"fmt"
	"math"
)

func main() {
	// F?1:T?4:5
	maze := [][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}
	fmt.Println(hasPath(maze, []int{0, 4}, []int{4, 4}))
	// fmt.Println(wordSquares([]string{"abat", "baba", "atan", "atal"}))
}

//lint:ignore U1000 unused
func powerTrue(k int) bool {
	for i := 0; i < 100; i++ {
		if k == int(math.Pow(2, float64(i))) {
			return true
		}
	}

	return false
}
