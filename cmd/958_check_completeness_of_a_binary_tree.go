package main

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}

	flag := false

	for len(queue) > 0 {
		newQ := []*TreeNode{}
		for _, node := range queue {
			if node == nil {
				flag = true
			} else {
				if flag {
					return false
				}
				newQ = append(newQ, node.Left)
				newQ = append(newQ, node.Right)
			}
		}
		queue = newQ
	}

	return true
}
