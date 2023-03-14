package main

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(0, root)
}

func sumNumbersHelper(prefix int, node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return prefix*10 + node.Val
	}

	left := sumNumbersHelper(prefix*10+node.Val, node.Left)
	right := sumNumbersHelper(prefix*10+node.Val, node.Right)

	return left + right
}
