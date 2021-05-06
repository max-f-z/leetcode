package main

func kthSmallest(root *TreeNode, k int) int {
	els := inOrder(root, []int{})

	return els[k-1]
}

func inOrder(node *TreeNode, els []int) []int {
	if node.Left != nil {
		els = inOrder(node.Left, els)
	}

	els = append(els, node.Val)

	if node.Right != nil {
		els = inOrder(node.Right, els)
	}

	return els
}
