package main

import (
	"fmt"
	"math"
)

func main() {
	root := &Node{Val: -1}
	// root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 1}
	// root.Left.Left = &Node{Val: 1}
	// root.Left.Right = &Node{Val: 3}
	root.Right.Right = &Node{Val: 9}

	fmt.Println(treeToDoublyList(root))
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
