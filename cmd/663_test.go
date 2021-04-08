package main

import (
	"testing"
)

func Test663(t *testing.T) {
	root3 := &TreeNode{
		Val: 0,
	}

	root := &TreeNode{
		Left: &TreeNode{
			Val: 10,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
			Val: 10,
		},
		Val: 5,
	}

	root2 := &TreeNode{
		Left: &TreeNode{
			Val: -1,
		},
		Right: &TreeNode{
			Val: 1,
		},
		Val: 0,
	}
	b := checkEqualTree(root3)
	t.Log(b)
	b = checkEqualTree(root)
	t.Log(b)
	b = checkEqualTree(root2)
	t.Log(b)
}
