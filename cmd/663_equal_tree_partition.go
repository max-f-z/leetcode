package main

var sums []int

func checkEqualTree(root *TreeNode) bool {
	sums = []int{}
	sum := checkEqualTreeHelper(root, true)

	for _, v := range sums {
		if v*2 == sum {
			return true
		}
	}

	return false
}

func checkEqualTreeHelper(node *TreeNode, isRoot bool) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		if !isRoot {
			sums = append(sums, node.Val)
		}

		return node.Val
	}

	var l, r int

	if node.Left != nil {
		l = checkEqualTreeHelper(node.Left, false)
	}

	if node.Right != nil {
		r = checkEqualTreeHelper(node.Right, false)
	}

	sum := l + r + node.Val

	if !isRoot {
		sums = append(sums, sum)
	}

	return sum
}
