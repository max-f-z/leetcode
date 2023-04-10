package main

import (
	"strings"
)

//lint:ignore U1000 unused
func intToRoman(num int) string {
	bits := []string{"I", "V", "X", "L", "C", "D", "M"}

	res := ""
	base := 0

	for num > 0 {
		tmp := num % 10

		var tmpStr strings.Builder
		switch tmp {
		case 0, 1, 2, 3:
			for i := 0; i < tmp; i++ {
				tmpStr.WriteString(bits[base])
			}
		case 4:
			tmpStr.WriteString(bits[base])
			tmpStr.WriteString(bits[base+1])
		case 5:
			tmpStr.WriteString(bits[base+1])
		case 6, 7, 8:
			tmpStr.WriteString(bits[base+1])
			for i := 0; i < tmp-5; i++ {
				tmpStr.WriteString(bits[base])
			}
		case 9:
			tmpStr.WriteString(bits[base])
			tmpStr.WriteString(bits[base+2])
		}
		res = tmpStr.String() + res

		num = num / 10
		base += 2
	}

	return res
}
