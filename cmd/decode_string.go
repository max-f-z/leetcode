package main

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	r, _ := decodeStr(s, 0)

	return r
}

func decodeStr(s string, start int) (res string, end int) {
	if len(s) == 0 {
		return "", 0
	}

	times := 0
	i := start
	for i = start; i < len(s); i++ {
		ch := string(s[i])
		if n, err := strconv.Atoi(ch); err == nil {
			times = 10*times + n
		} else if ch == "[" {
			decoded, end := decodeStr(s, i+1)
			res += strings.Repeat(decoded, times)
			i = end
			times = 0
		} else if ch == "]" {
			return res, i
		} else {
			res += ch
		}
	}

	return res, i
}
