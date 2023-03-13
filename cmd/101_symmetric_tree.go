package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirrorTree(root.Left, root.Right)
}

func mirrorTree(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	first := mirrorTree(l.Left, r.Right)
	second := mirrorTree(l.Right, r.Left)
	if first && second && l.Val == r.Val {
		return true
	}
	return false
}
