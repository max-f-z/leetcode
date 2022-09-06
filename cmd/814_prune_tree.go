package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if !pruneTreeHelper(root) {
		return nil
	}

	return root
}

func pruneTreeHelper(node *TreeNode) bool {
	if node == nil {
		return false
	}

	l := pruneTreeHelper(node.Left)
	r := pruneTreeHelper(node.Right)

	if !l {
		node.Left = nil
	}

	if !r {
		node.Right = nil
	}

	if node.Val == 1 || l || r {
		return true
	}

	return false
}
