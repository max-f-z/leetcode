package main

import "testing"

func Test783(t *testing.T) {
	root := &TreeNode{
		Val: 96,
		Left: &TreeNode{
			Val: 12,
			Right: &TreeNode{
				Val: 13,
				Right: &TreeNode{
					Val: 52,
					Left: &TreeNode{
						Val: 29,
					},
				},
			},
		},
	}

	ans := minDiffInBST(root)
	t.Log(ans)
}
