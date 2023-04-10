package main

//lint:ignore U1000 unused
func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, res, _ := countUnivalSubtreesHelper(root)

	return res
}

func countUnivalSubtreesHelper(root *TreeNode) (int, int, bool) {
	if root.Left == nil && root.Right == nil {
		return root.Val, 1, true
	}

	if root.Left != nil && root.Right == nil {
		value, cnt, sub := countUnivalSubtreesHelper(root.Left)
		if sub && value == root.Val {
			return value, cnt + 1, true
		}
		return value, cnt, false
	}

	if root.Right != nil && root.Left == nil {
		value, cnt, sub := countUnivalSubtreesHelper(root.Right)
		if sub && value == root.Val {
			return value, cnt + 1, true
		}
		return value, cnt, false
	}

	left, lcnt, lsub := countUnivalSubtreesHelper(root.Left)
	right, rcnt, rsub := countUnivalSubtreesHelper(root.Right)

	if lsub && rsub && left == right && left == root.Val {
		return root.Val, lcnt + rcnt + 1, true
	}

	return root.Val, lcnt + rcnt, false
}
