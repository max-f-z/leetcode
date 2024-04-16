package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//lint:ignore U1000 unused
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {

	root = addOneRowHelper(root, val, depth, 1, true)

	return root
}

func addOneRowHelper(node *TreeNode, val int, depth int, current int, left bool) *TreeNode {
	if current > depth {
		return node
	}

	if depth == current {
		cur := &TreeNode{
			Val: val,
		}

		if left {
			cur.Left = node
		} else {
			cur.Right = node
		}

		return cur
	}

	if node == nil {
		return node
	}

	node.Left = addOneRowHelper(node.Left, val, depth, current+1, true)
	node.Right = addOneRowHelper(node.Right, val, depth, current+1, false)

	return node
}
