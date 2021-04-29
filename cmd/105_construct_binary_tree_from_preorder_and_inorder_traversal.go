package main

var idx map[int]int
var pIdx int

func buildTree(preorder []int, inorder []int) *TreeNode {
	idx = map[int]int{}
	pIdx = 0
	for i := 0; i < len(inorder); i++ {
		idx[inorder[i]] = i
	}

	return buildTreeHelper(0, len(preorder)-1, preorder)
}

func buildTreeHelper(left, right int, preorder []int) *TreeNode {
	if left > right {
		return nil
	}

	node := &TreeNode{
		Val: preorder[pIdx],
	}

	pIdx++

	node.Left = buildTreeHelper(left, idx[node.Val]-1, preorder)

	node.Right = buildTreeHelper(idx[node.Val]+1, right, preorder)

	return node
}
