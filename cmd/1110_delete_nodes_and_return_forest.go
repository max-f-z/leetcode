package main

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ansMap := make(map[*TreeNode]bool)

	toDeleteMap := map[int]bool{}
	for i := 0; i < len(to_delete); i++ {
		toDeleteMap[to_delete[i]] = true
	}

	ansMap[root] = true

	delNodesHelper(root, toDeleteMap, ansMap, nil)

	ans := make([]*TreeNode, len(ansMap))

	idx := 0
	for k := range ansMap {
		ans[idx] = k
		idx++
	}

	return ans
}

func delNodesHelper(node *TreeNode, ref map[int]bool, ansMap map[*TreeNode]bool, parent *TreeNode) {
	if node == nil {
		return
	}

	if ref[node.Val] {
		if parent != nil {
			if parent.Left == node {
				parent.Left = nil
			}
			if parent.Right == node {
				parent.Right = nil
			}
		}
		delete(ansMap, node)

		if node.Left != nil {
			ansMap[node.Left] = true
		}

		if node.Right != nil {
			ansMap[node.Right] = true
		}
	}
	delNodesHelper(node.Left, ref, ansMap, node)
	delNodesHelper(node.Right, ref, ansMap, node)
}
