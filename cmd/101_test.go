package main

import "testing"

func Test101(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	ans := isSymmetric(root)

	t.Log(ans)
}
