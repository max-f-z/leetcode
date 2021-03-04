package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestorHelper(root, p, q)
}

func lowestCommonAncestorHelper(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := lowestCommonAncestorHelper(root.Left, p, q)
	right := lowestCommonAncestorHelper(root.Right, p, q)

	if root.Val == p.Val || root.Val == q.Val || (left != nil && right != nil) {
		return root
	}

	if left == nil {
		return right
	}

	return left
}
