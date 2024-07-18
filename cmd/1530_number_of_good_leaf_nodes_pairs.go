package main

func countPairs(root *TreeNode, distance int) int {
	ans := 0
	countPairsHelper(root, distance, 0, &ans)
	return ans
}

func countPairsHelper(node *TreeNode, distance int, level int, ans *int) []int {
	if node == nil {
		return []int{}
	}

	if node.Left == nil && node.Right == nil {
		return []int{level}
	}

	left := countPairsHelper(node.Left, distance, level+1, ans)
	right := countPairsHelper(node.Right, distance, level+1, ans)

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			if left[i]-level+right[j]-level <= distance {
				*ans += 1
			}
		}
	}

	return append(left, right...)
}
