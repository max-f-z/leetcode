package main

import "strings"

//lint:ignore U1000 unused
func reverseWords151(s string) string {
	strs := strings.Split(s, " ")

	ans := strings.Builder{}

	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != "" {
			ans.WriteString(strs[i])
			ans.WriteString(" ")
		}
	}

	return strings.TrimSpace(ans.String())
}
