package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1530(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	ans := countPairs(root, 3)
	assert.Check(t, ans == 1)
}
