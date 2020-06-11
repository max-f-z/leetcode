package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	bucket := []*strings.Builder{}
	for i := 0; i < numRows; i++ {
		bucket = append(bucket, &strings.Builder{})
	}

	index := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		bucket[index].WriteByte(s[i])
		if sign == 1 {
			index++
			if index >= numRows {
				index = index - 2
				sign = -1
			}
		} else {
			index--
			if index < 0 {
				index = index + 2
				sign = 1
			}
		}
	}

	res := ""
	for i := 0; i < numRows; i++ {
		res += bucket[i].String()
	}

	return res
}
