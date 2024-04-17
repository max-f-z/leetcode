package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	return smallestFromLeafHelper(root, "")
}

func smallestFromLeafHelper(node *TreeNode, suffix string) string {
	if node == nil {
		return suffix
	}

	suffix = string(byte('a'+node.Val)) + suffix

	l := smallestFromLeafHelper(node.Left, suffix)
	r := smallestFromLeafHelper(node.Right, suffix)

	ans := string(byte('a' + 27))

	if node.Left != nil && l < ans {
		ans = l
	}

	if node.Right != nil && r < ans {
		ans = r
	}

	if ans == string(byte('a'+27)) {
		return suffix
	}

	return ans
}
