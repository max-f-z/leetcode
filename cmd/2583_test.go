package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2583(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	ans := kthLargestLevelSum(root, 1)

	assert.Check(t, ans == 3)
}
