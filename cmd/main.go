package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rearrangeString("a", 0))
}

func powerTrue(k int) bool {
	for i := 0; i < 100; i++ {
		if k == int(math.Pow(2, float64(i))) {
			return true
		}
	}

	return false
}
