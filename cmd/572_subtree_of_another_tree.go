package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return isSubTreeHelper(s, t)
}

func isSubTreeHelper(node *TreeNode, t *TreeNode) bool {
	if isIdenticalTree(node, t) {
		return true
	}

	if node.Left != nil && isSubTreeHelper(node.Left, t) {
		return true
	}

	if node.Right != nil && isSubTreeHelper(node.Right, t) {
		return true
	}

	return false
}

func isIdenticalTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val == t2.Val {
		if t1.Left == nil && t1.Right == nil && t2.Right == nil && t2.Left == nil {
			return true
		}

		return isIdenticalTree(t1.Left, t2.Left) && isIdenticalTree(t1.Right, t2.Right)
	}

	return false
}
