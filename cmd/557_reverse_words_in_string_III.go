package main

import "strings"

func reverseWords557(s string) string {
	ss := strings.Split(s, " ")

	sb := strings.Builder{}
	for idx, str := range ss {
		for i := len(str) - 1; i >= 0; i-- {
			sb.WriteByte(str[i])
		}

		if idx != len(ss)-1 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}
