package main

import (
	"strconv"
	"strings"
)

func tree2str(root *TreeNode) string {
	return preorderHelper(root)
}

func preorderHelper(node *TreeNode) string {
	sb := strings.Builder{}
	if node == nil {
		return sb.String()
	}
	sb.WriteString(strconv.Itoa(node.Val))

	l := preorderHelper(node.Left)
	r := preorderHelper(node.Right)

	if l != "" || r != "" {
		sb.WriteRune('(')
		sb.WriteString(l)
		sb.WriteRune(')')
	}

	if r != "" {
		sb.WriteRune('(')
		sb.WriteString(r)
		sb.WriteRune(')')
	}

	return sb.String()
}
