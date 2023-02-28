package main

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	mem := map[string]int{}
	ans := []*TreeNode{}
	findDuplicateSubtreesHelper(root, mem, &ans)
	return ans
}

func findDuplicateSubtreesHelper(node *TreeNode, mem map[string]int, ans *[]*TreeNode) string {
	if node == nil {
		return ""
	}

	l := findDuplicateSubtreesHelper(node.Left, mem, ans)
	r := findDuplicateSubtreesHelper(node.Right, mem, ans)

	str := fmt.Sprintf("(%s)(%d)(%s)", l, node.Val, r)
	if str != "" {
		mem[str]++
	}
	if mem[str] == 2 {
		(*ans) = append((*ans), node)
	}

	return str
}
