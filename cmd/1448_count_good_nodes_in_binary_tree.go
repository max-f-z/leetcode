package main

func goodNodes(root *TreeNode) int {
	return goodNodesHelper(root, root.Val)
}

func goodNodesHelper(node *TreeNode, max int) int {
	ans := 0
	if node.Val >= max {
		ans++
	}

	local := max
	if node.Val > max {
		local = node.Val
	}

	if node.Left != nil {

		l := goodNodesHelper(node.Left, local)
		ans += l
	}

	if node.Right != nil {
		r := goodNodesHelper(node.Right, local)
		ans += r

	}

	return ans
}
