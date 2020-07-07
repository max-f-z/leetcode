package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sortTransformedArray([]int{-4, -2, 2, 4}, 0, -1, 5))
}

func powerTrue(k int) bool {
	for i := 0; i < 100; i++ {
		if k == int(math.Pow(2, float64(i))) {
			return true
		}
	}

	return false
}
