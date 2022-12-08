package main

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := []int{}
	leaf2 := []int{}

	getLeaves(root1, &leaf1)
	getLeaves(root2, &leaf2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}

	return true
}

func getLeaves(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
		return
	}

	if node.Left != nil {
		getLeaves(node.Left, leaves)
	}

	if node.Right != nil {
		getLeaves(node.Right, leaves)
	}
}
