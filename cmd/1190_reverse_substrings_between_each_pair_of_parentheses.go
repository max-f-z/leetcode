package main

import "strings"

func reverseParentheses(s string) string {
	stack := []string{}

	tmp := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, tmp.String())
			tmp = strings.Builder{}
		} else if s[i] == ')' {
			rev := reverseString(tmp.String())
			tmp = strings.Builder{}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp.WriteString(pop)
			tmp.WriteString(rev)
		} else {
			tmp.WriteByte(s[i])
		}
	}

	return tmp.String()
}
