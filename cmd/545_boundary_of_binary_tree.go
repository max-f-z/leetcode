package main

var ans545 []int

//lint:ignore U1000 unused
func boundaryOfBinaryTree(root *TreeNode) []int {
	ans545 = []int{}

	boundaryOfBinaryTreeDFS(root.Left, 1)
	boundaryOfBinaryTreeDFS(root.Right, -1)

	return ans545
}

func boundaryOfBinaryTreeDFS(node *TreeNode, flag int) {
	if node == nil {
		return
	}

	if flag == 1 { // Left
		ans545 = append(ans545, node.Val)

		if node.Left != nil {
			boundaryOfBinaryTreeDFS(node.Left, 1)
			boundaryOfBinaryTreeDFS(node.Right, 0)
		} else {
			boundaryOfBinaryTreeDFS(node.Right, 1)
		}

	} else if flag == -1 { // right
		if node.Right != nil {
			boundaryOfBinaryTreeDFS(node.Left, 0)
			boundaryOfBinaryTreeDFS(node.Right, -1)
		} else {
			boundaryOfBinaryTreeDFS(node.Left, -1)
		}
		ans545 = append(ans545, node.Val)
	} else { // find leafs
		if node.Left == nil && node.Right == nil {
			ans545 = append(ans545, node.Val)
		}
		boundaryOfBinaryTreeDFS(node.Left, 0)
		boundaryOfBinaryTreeDFS(node.Right, 0)
	}

}
