package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test988(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	str := smallestFromLeaf(root)
	assert.Check(t, str == "dba")

	root = &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
		},
	}

	str = smallestFromLeaf(root)
	assert.Check(t, str == "ba")
}
