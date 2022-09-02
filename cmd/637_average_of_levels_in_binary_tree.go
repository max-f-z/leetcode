package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	results := []float64{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		newQ := []*TreeNode{}

		sum := 0
		for _, node := range queue {
			sum += node.Val
			if node.Left != nil {
				newQ = append(newQ, node.Left)
			}
			if node.Right != nil {
				newQ = append(newQ, node.Right)
			}
		}

		results = append(results, float64(sum)/float64(len(queue)))
		queue = newQ
	}

	return results
}
