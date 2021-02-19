package main

import "math"

var maxVal int

func longestConsecutive542(root *TreeNode) int {
	maxVal = 0
	longestConsecutiveDFS(root)
	return maxVal
}

func longestConsecutiveDFS(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	inc := 1
	dec := 1

	l := longestConsecutiveDFS(root.Left)
	r := longestConsecutiveDFS(root.Right)

	if root.Left != nil {
		if root.Left.Val == root.Val-1 {
			dec = int(math.Max(float64(dec), float64(l[1]+1)))
		}

		if root.Left.Val == root.Val+1 {
			inc = int(math.Max(float64(inc), float64(l[0]+1)))
		}
	}

	if root.Right != nil {
		if root.Right.Val == root.Val+1 {
			inc = int(math.Max(float64(inc), float64(r[0]+1)))
		}

		if root.Right.Val == root.Val-1 {
			dec = int(math.Max(float64(dec), float64(r[1]+1)))
		}
	}

	maxVal = int(math.Max(float64(maxVal), float64(dec+inc-1)))

	return []int{inc, dec}
}
