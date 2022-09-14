package main

import "testing"

// [2,1,1,1,3,null,null,null,null,null,1]
// [2,3,1,3,1,null,1]
func Test1457(t *testing.T) {
	node := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	ans := pseudoPalindromicPaths(node)

	t.Log(ans)
}
