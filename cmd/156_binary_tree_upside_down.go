package main

//lint:ignore U1000 unused
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	newRoot := upsideDownBinaryTree(root.Left)
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Right = nil
	root.Left = nil
	return newRoot
}
