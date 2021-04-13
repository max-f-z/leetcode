package main

import (
	"container/list"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	remove := []int{}
	stack := list.New()

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.PushBack(i)
		}

		if s[i] == ')' {
			if stack.Len() > 0 {
				stack.Remove(stack.Back())
			} else {
				remove = append(remove, i)
			}
		}
	}

	nums := stack.Len()
	for i := 0; i < nums; i++ {
		tmp := stack.Front().Value.(int)
		stack.Remove(stack.Front())
		remove = append(remove, tmp)
	}

	sb := strings.Builder{}

	for i := 0; i < len(s); i++ {
		match := false
		for j := 0; j < len(remove); j++ {
			if i == remove[j] {
				match = true
				break
			}
		}

		if !match {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}
