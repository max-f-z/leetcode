package main

import "math"

var val int

//lint:ignore U1000 unused
func diameterOfBinaryTree(root *TreeNode) int {
	val = math.MinInt32

	tmp := diameterOfBinaryTreeHelper(root)

	if val < tmp {
		val = tmp
	}

	return val - 1
}

func diameterOfBinaryTreeHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	l := diameterOfBinaryTreeHelper(node.Left)
	r := diameterOfBinaryTreeHelper(node.Right)

	if val < l+r+1 {
		val = l + r + 1
	}

	if l > r {
		return l + 1
	}
	return r + 1
}
