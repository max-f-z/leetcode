package main

import "strings"

//lint:ignore U1000 unused
func fullJustify(words []string, maxWidth int) []string {
	res := []string{}

	idx := 0
	tmp := ""
	for idx < len(words) {
		idx, tmp = fullJustifyHelper(words, idx, maxWidth)
		res = append(res, tmp)
	}

	return res
}

func fullJustifyHelper(words []string, start int, maxWidth int) (int, string) {
	res := []string{}
	sumL := 0
	idx := start

	for i := idx; i < len(words); i++ {
		if len(res) == 0 {
			res = append(res, words[i])
			sumL += len(words[i])
			continue
		}

		if len(res)+sumL+len(words[i]) <= maxWidth {
			res = append(res, words[i])
			sumL += len(words[i])
			idx = i
			continue
		}

		if len(res)+sumL+len(words[i]) > maxWidth {
			break
		}
	}

	sb := strings.Builder{}
	if idx != len(words)-1 {
		var breaks []string

		if len(res) > 1 {
			breaks = make([]string, len(res)-1)
		} else {
			breaks = make([]string, 1)
		}

		cursor := 0
		for i := sumL; i < maxWidth; i++ {
			if cursor < len(res)-1 {
				breaks[cursor] = breaks[cursor] + " "
				cursor++
				continue
			}

			cursor = 0
			breaks[cursor] = breaks[cursor] + " "
			cursor++
		}

		for i := 0; i < len(res); i++ {
			sb.WriteString(res[i])
			if i < len(res)-1 {
				sb.WriteString(breaks[i])
			}
		}

		if len(res) == 1 {
			sb.WriteString(breaks[0])
		}
	} else {
		for i := 0; i < len(res); i++ {
			sb.WriteString(res[i])
			if sumL+1 <= maxWidth {
				sumL++
				sb.WriteString(" ")
			}
		}

		for i := sumL; i < maxWidth; i++ {
			sb.WriteString(" ")
		}
	}

	return idx + 1, sb.String()
}
