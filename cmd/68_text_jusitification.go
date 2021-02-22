package main

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}

	idx := 0
	tmp := ""

	for idx < len(words)+1 {
		idx, tmp = fullJustifyHelper(words, maxWidth, idx)
		res = append(res, tmp)
	}

	return res
}

func fullJustifyHelper(words []string, maxWidth int, start int) (int, string) {
	l := 0
	strB := strings.Builder{}
	idx := start
	breaks := []string{}
	remain := false

	for idx = start; idx < len(words); idx++ {
		if l+len(words[idx])+1 <= maxWidth+1 {
			l += len(words[idx]) + 1
			breaks = append(breaks, " ")
			continue
		}
		remain = true
		break
	}

	if !remain {
		l = 0
		for i := start; i < len(words); i++ {
			strB.WriteString(words[i])
			l += len(words[i])
			if l+1 <= maxWidth {
				strB.WriteString(" ")
				l++
			}
		}

		for l < maxWidth {
			strB.WriteString(" ")
			l++
		}
		return len(words) + 1, strB.String()
	}

	l--
	breaks[len(breaks)-1] = ""
	if len(breaks) == 1 {
		strB.WriteString(words[start])
		for i := 0; i < maxWidth-len(words[start]); i++ {
			strB.WriteString(" ")
		}
	} else {
		cursor := 0
		for (l + 1) <= maxWidth {
			if cursor == len(breaks)-1 {
				cursor = 0
			}
			breaks[cursor] = breaks[cursor] + " "
			l++
			cursor++

		}

		for i := 0; i < len(breaks); i++ {
			strB.WriteString(words[start+i])
			strB.WriteString(breaks[i])
		}
	}
	return idx, strB.String()
}
