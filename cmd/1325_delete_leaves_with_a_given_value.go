package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	ans := removeLeafNodesHelper(root, target)

	return ans
}

func removeLeafNodesHelper(node *TreeNode, target int) *TreeNode {
	if node == nil {
		return nil
	}

	l := removeLeafNodesHelper(node.Left, target)
	r := removeLeafNodesHelper(node.Right, target)

	node.Left = l
	node.Right = r

	if l == nil && r == nil && node.Val == target {
		return nil
	}

	return node
}
