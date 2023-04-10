package main

import (
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
func str2tree(s string) *TreeNode {
	if s == "()" || s == "" {
		return nil
	}

	idxs := strings.Index(s, "(")
	var val int
	node := &TreeNode{}
	if idxs != -1 {
		val, _ = strconv.Atoi(s[:idxs])
		node.Val = val
	} else {
		val, _ = strconv.Atoi(s)
		node.Val = val
		return node
	}

	s = s[idxs:]
	count := 0

	var leftStr, rightStr string

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		}

		if s[i] == ')' {
			count--
			if count == 0 {
				leftStr = s[:i+1]
				rightStr = s[i+1:]
				break
			}
		}
	}

	if leftStr == "()" || leftStr == "" {
		node.Left = nil
	} else {
		node.Left = str2tree(leftStr[1 : len(leftStr)-1])
	}

	if rightStr == "()" || rightStr == "" {
		node.Right = nil
	} else {
		node.Right = str2tree(rightStr[1 : len(rightStr)-1])
	}

	return node
}
