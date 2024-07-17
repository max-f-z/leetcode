package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1110(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	ans := delNodes(root, []int{3, 5})

	assert.Check(t, len(ans) == 3)
}
