package main

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	n1, n2 := make([]int, l1), make([]int, l2)

	for i := 0; i < l1; i++ {
		val, _ := strconv.Atoi(string(num1[i]))
		n1[l1-1-i] = val
	}

	for i := 0; i < l2; i++ {
		val, _ := strconv.Atoi(string(num2[i]))
		n2[l2-1-i] = val
	}

	s := l1
	if s < l2 {
		s = l2
	}
	sum := make([]int, s)
	for i := 0; i < s; i++ {
		if i < l1 {
			sum[i] += n1[i]
		}
		if i < l2 {
			sum[i] += n2[i]
		}
	}

	var sb strings.Builder

	tmp := 0
	for i := 0; i < s; i++ {
		sum[i] += tmp
		tmp = 0
		if sum[i] >= 10 {
			sum[i] = sum[i] - 10
			tmp = 1
		}

		sb.WriteString(strconv.Itoa(sum[i]))
	}

	if tmp == 1 {
		sb.WriteString("1")
	}

	return Reverse(sb.String())
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
