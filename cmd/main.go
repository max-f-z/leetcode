package main

import (
	"fmt"
	"math"
)

func main() {
	node5 := &Node428{Val: 5}
	node6 := &Node428{Val: 6}
	node3 := &Node428{Val: 3, Children: []*Node428{node5, node6}}
	node2 := &Node428{Val: 2}
	node4 := &Node428{Val: 4}
	root := &Node428{Val: 1, Children: []*Node428{node3, node2, node4}}
	c := Constructor428()
	tmp := c.serialize(root)
	fmt.Println(tmp)
	newNode := c.deserialize(tmp)
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
