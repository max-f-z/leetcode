package main

import (
	"container/list"
	"strings"
)

func removeDuplicates1047(s string) string {
	stack := list.New()

	for i := 0; i < len(s); i++ {
		if stack.Len() == 0 {
			stack.PushBack(s[i])
			continue
		}

		if stack.Back().Value.(byte) == s[i] {
			stack.Remove(stack.Back())
			continue
		}

		stack.PushBack(s[i])
	}

	sb := strings.Builder{}

	for stack.Front() != nil {
		sb.WriteByte(stack.Front().Value.(byte))
		stack.Remove(stack.Front())
	}

	return sb.String()
}
