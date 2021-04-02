package main

import (
	"strings"
)

func SpinWords(str string) string {
	strs := strings.Split(str, " ")
	sb := strings.Builder{}
	for _, v := range strs {
		if len(v) >= 5 {
			sb.WriteString(reverseString(v))
		} else {
			sb.WriteString(v)
		}
		sb.WriteString(" ")
	}
	l := len(sb.String())
	res := sb.String()[:l-1]
	return res
} // SpinWords

func reverseString(str string) string {
	sb := strings.Builder{}

	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}

	return sb.String()
}
