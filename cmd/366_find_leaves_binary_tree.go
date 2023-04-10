package main

//lint:ignore U1000 unused
func findLeaves(root *TreeNode) [][]int {
	res := [][]int{}
	for root != nil {
		tmp := []int{}

		if findLeavesHelper(root, nil, &tmp) {
			res = append(res, tmp)
			break
		}

		res = append(res, tmp)
	}
	return res
}

func findLeavesHelper(node, p *TreeNode, res *[]int) bool {
	if node.Left == nil && node.Right == nil {
		*res = append(*res, node.Val)
		if p != nil && p.Left == node {
			p.Left = nil
		}

		if p != nil && p.Right == node {
			p.Right = nil
		}
		return true
	}

	if node.Left != nil {
		findLeavesHelper(node.Left, node, res)
	}

	if node.Right != nil {
		findLeavesHelper(node.Right, node, res)
	}
	return false
}
