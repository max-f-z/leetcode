package main

//lint:ignore U1000 unused
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val && root.Val > q.Val {
		return root
	}

	if root.Val > p.Val && root.Val < q.Val {
		return root
	}

	if root.Val == p.Val {
		return root
	}

	if root.Val == q.Val {
		return root
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor235(root.Right, p, q)
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor235(root.Left, p, q)
	}

	return nil
}
