package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	d := &dfs783{
		root: root,
		min:  math.MaxInt32,
	}

	d.minDiffInBSTHelper(root)
	return d.min
}

type dfs783 struct {
	root *TreeNode
	min  int
}

func (d *dfs783) minDiffInBSTHelper(node *TreeNode) (int, int) {
	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val
	}

	l, r := node.Val, node.Val
	if node.Left != nil {
		ll, lr := d.minDiffInBSTHelper(node.Left)
		if node.Val-lr < d.min {
			d.min = node.Val - lr
		}
		if ll < l {
			l = ll
		}
	}

	if node.Right != nil {
		rl, rr := d.minDiffInBSTHelper(node.Right)
		if rl-node.Val < d.min {
			d.min = rl - node.Val
		}
		if rr > r {
			r = rr
		}
	}
	return l, r
}
