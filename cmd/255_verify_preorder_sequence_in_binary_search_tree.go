package main

import "math"

//lint:ignore U1000 unused
func verifyPreorder(preorder []int) bool {
	return verifyPreorderHelper(preorder, 0, len(preorder)-1, math.MinInt32, math.MaxInt32)
}

func verifyPreorderHelper(preorder []int, start, end int, min, max int) bool {
	if start > end {
		return true
	}

	root := preorder[start]

	idx := start

	for idx = start; idx <= end; idx++ {
		val := preorder[idx]

		if val > max || val < min {
			return false
		}

		if val > root {
			break
		}
	}

	left := verifyPreorderHelper(preorder, start+1, idx-1, min, root)
	right := verifyPreorderHelper(preorder, idx, end, root, max)

	if left && right {
		return true
	}
	return false
}
