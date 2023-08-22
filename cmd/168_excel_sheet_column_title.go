package main

import "strings"

func convertToTitle(columnNumber int) string {
	sb := strings.Builder{}

	for columnNumber != 0 {
		val := columnNumber % 26

		if val == 0 {
			sb.WriteByte('Z')
			columnNumber -= 26
		} else {
			sb.WriteByte(byte('A' + val - 1))
		}
		columnNumber /= 26
	}

	return reverseString(sb.String())
}

// func reverseString(str string) string {
// 	bytes := []byte(str)
// 	n := len(str)
// 	for i := 0; i < n/2; i++ {
// 		bytes[i], bytes[n-i-1] = bytes[n-i-1], bytes[i]
// 	}
// 	return string(bytes)
// }
