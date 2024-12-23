package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations2471(root *TreeNode) int {
	next := []*TreeNode{root}

	ops := 0

	for len(next) > 0 {
		new_next := []*TreeNode{}
		level := make([]int, len(next))
		sortedLevel := make([]int, len(next))
		for idx, node := range next {
			if node.Left != nil {
				new_next = append(new_next, node.Left)
			}

			if node.Right != nil {
				new_next = append(new_next, node.Right)
			}
			level[idx] = node.Val
		}

		copy(sortedLevel, level)
		sort.Ints(sortedLevel)

		mapSortedLevel := map[int]int{}

		for i := 0; i < len(next); i++ {
			mapSortedLevel[sortedLevel[i]] = i
		}

		levelOps := 0
		for i := 0; i < len(next); i++ {
			if level[i] != sortedLevel[i] {
				shouldIdx := mapSortedLevel[level[i]]

				level[i], level[shouldIdx] = level[shouldIdx], level[i]
				levelOps += 1
				i--
			}
		}

		ops += levelOps
		next = new_next
	}

	return ops
}
