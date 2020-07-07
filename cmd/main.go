package main

import (
	"fmt"
	"math"
)

func main() {
	obj := Constructor362()
	obj.Hit(1)
	obj.Hit(2)
	obj.Hit(3)

	fmt.Println(obj.GetHits((4)))
	obj.Hit(300)
	fmt.Println(obj.GetHits(300))
	fmt.Println(obj.GetHits(301))
}

func powerTrue(k int) bool {
	for i := 0; i < 100; i++ {
		if k == int(math.Pow(2, float64(i))) {
			return true
		}
	}

	return false
}
