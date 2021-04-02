package main

import (
	"strings"
)

func splitLoopedString(strs []string) string {

	for i := 0; i < len(strs); i++ {
		reverse := reverseString(strs[i])
		if strs[i] < reverse {
			strs[i] = reverse
		}
	}

	max := ""

	for i := 0; i < len(strs); i++ {
		rev := reverseString(strs[i])
		for j := 0; j < len(strs[i]); j++ {
			sb := strings.Builder{}
			rsb := strings.Builder{}
			sb.WriteString(strs[i][j:])
			rsb.WriteString(rev[j:])

			for k := i + 1; k < len(strs); k++ {
				sb.WriteString(strs[k])
				rsb.WriteString(strs[k])
			}

			for k := 0; k < i; k++ {
				sb.WriteString(strs[k])
				rsb.WriteString(strs[k])
			}

			sb.WriteString(strs[i][0:j])
			rsb.WriteString(rev[0:j])

			if sb.String() > max {
				max = sb.String()
			}

			if rsb.String() > max {
				max = rsb.String()
			}
		}
	}

	return max
}
