package main

import (
	"strings"
	"unicode"
)

func makeGood(s string) string {
	if len(s) <= 1 {
		return s
	}

	sb := strings.Builder{}

	flag := false

	i := 0

	for ; i < len(s)-1; i++ {
		upper := unicode.ToUpper(rune(s[i]))
		nextUpper := unicode.ToUpper(rune(s[i+1]))

		if s[i] != s[i+1] && upper == nextUpper {
			i++
			flag = true
			continue
		} else {
			sb.WriteByte(s[i])
		}
	}
	if i < len(s) {
		sb.WriteByte(s[i])
	}

	if !flag {
		return sb.String()
	}

	return makeGood(sb.String())
}
