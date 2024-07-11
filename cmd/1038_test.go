package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1038(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}

	original := BSTtoList(root)
	assert.DeepEqual(t, original, []int{4, 1, 0, 2, 3, 6, 5, 7, 8})

	bstToGst(root)
	ans := BSTtoList(root)

	assert.DeepEqual(t, ans, []int{30, 36, 36, 35, 33, 21, 26, 15, 8})
}

func BSTtoList(node *TreeNode) []int {
	// preorder
	ans := []int{}
	BSTtoListHelp(node, &ans)
	return ans
}

func BSTtoListHelp(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	*ans = append(*ans, node.Val)
	BSTtoListHelp(node.Left, ans)
	BSTtoListHelp(node.Right, ans)
}
