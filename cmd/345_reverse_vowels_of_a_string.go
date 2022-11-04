package main

import "strings"

func reverseVowels(s string) string {
	p1, p2 := 0, len(s)-1
	sb := strings.Builder{}

	for ; p1 < len(s); p1++ {
		if isVowel(s[p1]) {
			for ; p2 >= 0; p2-- {
				if isVowel(s[p2]) {
					sb.WriteByte(s[p2])
					p2--
					break
				}
			}
		} else {
			sb.WriteByte(s[p1])
		}
	}

	return sb.String()
}

func isVowel(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
