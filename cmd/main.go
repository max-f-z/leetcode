package main

import "fmt"

func main() {
	// fmt.Println(alienOrder([]string{"bsusz", "rhn", "gfbrwec", "kuw", "qvpxbexnhx", "wnp", "laxutz", "czxccww"}))
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	a := []int{3, 5, 6, 7, 8, 9}
	wiggleSort(a)
	fmt.Println(a)
}
