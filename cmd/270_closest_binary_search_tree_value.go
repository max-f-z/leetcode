package main

import "math"

func closestValue(root *TreeNode, target float64) int {
	res := closestValueHelper(root, target, math.MaxInt64)
	return res
}

func closestValueHelper(node *TreeNode, target float64, res int) int {
	if node == nil {
		return res
	}

	tmp := math.Abs(float64(node.Val) - target)
	if tmp < math.Abs(float64(res)-target) {
		res = node.Val
	}

	res = closestValueHelper(node.Left, target, res)
	res = closestValueHelper(node.Right, target, res)

	return res
}
