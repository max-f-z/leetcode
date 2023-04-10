package main

//lint:ignore U1000 unused
func rangeSumBST(root *TreeNode, low int, high int) int {
	return rangeSumBSTHelper(root, low, high)
}

func rangeSumBSTHelper(current *TreeNode, low int, high int) int {
	if current == nil {
		return 0
	}

	sum := 0

	if current.Val >= low && current.Val <= high {
		sum += current.Val
	}

	if current.Left != nil {
		sum += rangeSumBSTHelper(current.Left, low, high)
	}

	if current.Right != nil {
		sum += rangeSumBSTHelper(current.Right, low, high)
	}

	return sum
}
