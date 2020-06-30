package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(alienOrder([]string{"bsusz", "rhn", "gfbrwec", "kuw", "qvpxbexnhx", "wnp", "laxutz", "czxccww"}))
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}

	fmt.Println(inorderSuccessor(root, root.Right))

	wallsAndGates([][]int{{math.MaxInt32, -1, 0, math.MaxInt32}, {math.MaxInt32, math.MaxInt32, math.MaxInt32, -1}, {math.MaxInt32, -1, math.MaxInt32, -1}, {0, -1, math.MaxInt32, math.MaxInt32}})
}
