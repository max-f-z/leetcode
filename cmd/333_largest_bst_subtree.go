package main

import "math"

//lint:ignore U1000 unused
func largestBSTSubtree(root *TreeNode) int {
	res := 0
	largestBSTSubtreeHelper(&res, root)
	return res
}

func largestBSTSubtreeHelper(res *int, node *TreeNode) (bool, int64, int64, int) {
	if node == nil {
		return true, math.MaxInt64, math.MinInt64, 0
	}

	isLeftBST, leftMn, leftMx, leftCount := largestBSTSubtreeHelper(res, node.Left)
	isRightBST, rightMn, rightMx, rightCount := largestBSTSubtreeHelper(res, node.Right)
	count := leftCount + rightCount + 1
	if !isLeftBST || !isRightBST || leftMx >= int64(node.Val) || rightMn <= int64(node.Val) {
		return false, math.MinInt64, math.MaxInt64, count
	}

	if count > *res {
		*res = count
	}
	return true, min2(int64(node.Val), leftMn), max(int64(node.Val), rightMx), count
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}

	return y
}

func min2(x, y int64) int64 {
	if x < y {
		return x
	}

	return y
}
