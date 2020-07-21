package main

import (
	"fmt"
	"math"
)

func main() {
	// F?1:T?4:5
	fmt.Println(minTransfers([][]int{{2, 0, 5}, {4, 3, 4}}))
	// fmt.Println(wordSquares([]string{"abat", "baba", "atan", "atal"}))
}

func powerTrue(k int) bool {
	for i := 0; i < 100; i++ {
		if k == int(math.Pow(2, float64(i))) {
			return true
		}
	}

	return false
}
