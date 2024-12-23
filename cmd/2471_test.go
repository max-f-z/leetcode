package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2471(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 10,
				},
			},
		},
	}

	ans := minimumOperations2471(root)
	assert.Check(t, ans == 3)
}
