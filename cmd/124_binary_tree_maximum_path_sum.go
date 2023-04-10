package main

import "math"

var res int

//lint:ignore U1000 unused
func maxPathSum(root *TreeNode) int {
	res = int(math.MinInt32)

	maxPathSumHelper(root)

	return res
}

func maxPathSumHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := maxPathSumHelper(node.Left)
	if l < 0 {
		l = 0
	}

	r := maxPathSumHelper(node.Right)
	if r < 0 {
		r = 0
	}

	sum := l + r + node.Val
	if sum > res {
		res = sum
	}

	if l < r {
		return r + node.Val
	}
	return l + node.Val
}
