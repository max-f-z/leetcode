package main

import (
	"fmt"
	"math"
)

func main() {
	node5 := &Node431{Val: 5}
	node6 := &Node431{Val: 6}
	node3 := &Node431{Val: 3, Children: []*Node431{node5, node6}}
	node2 := &Node431{Val: 2}
	node4 := &Node431{Val: 4}
	root := &Node431{Val: 1, Children: []*Node431{node3, node2, node4}}
	c := Constructor431()
	tmp := c.encode(root)
	fmt.Println(tmp)
	newNode := c.decode(tmp)
	fmt.Println(newNode)
	fmt.Println(tmp)

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
