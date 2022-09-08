package main

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	l := inorderTraversal(root.Left)

	res = append(res, l...)
	res = append(res, root.Val)

	r := inorderTraversal(root.Right)

	res = append(res, r...)

	return res
}
