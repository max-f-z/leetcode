package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node429 struct {
	Val      int
	Children []*Node429
}

func levelOrder(root *Node429) [][]int {
	queue := []*Node429{root}
	results := [][]int{}
	if root == nil {
		return results
	}

	for len(queue) > 0 {
		newQ := []*Node429{}
		level := []int{}
		for _, node := range queue {
			if node.Children != nil && len(node.Children) > 0 {
				for _, child := range node.Children {
					newQ = append(newQ, child)
				}
			}
			level = append(level, node.Val)
		}

		results = append(results, level)
		queue = newQ
	}

	return results
}
