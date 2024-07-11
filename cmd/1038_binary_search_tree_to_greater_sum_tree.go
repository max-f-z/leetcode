package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}

	bstToGstHelper(root, &stack)

	sum := 0

	for i := 0; i < len(stack); i++ {
		sum += stack[i].Val
		stack[i].Val = sum
	}

	return root
}

func bstToGstHelper(node *TreeNode, stack *[]*TreeNode) {
	if node == nil {
		return
	}

	bstToGstHelper(node.Right, stack)
	*stack = append(*stack, node)
	bstToGstHelper(node.Left, stack)
}
