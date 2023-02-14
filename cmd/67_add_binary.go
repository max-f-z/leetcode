package main

import "strings"

func addBinary(a string, b string) string {
	if a == "0" {
		return b
	}

	if b == "0" {
		return a
	}

	m, n := len(a), len(b)

	op1 := strings.Split(a, "")
	op2 := strings.Split(b, "")
	reverseStringArray(&op1)
	reverseStringArray(&op2)

	ans := []string{}

	carry := 0

	for p1, p2 := 0, 0; p1 < m || p2 < n; {
		sum := 0
		sum += carry
		if p1 < m && op1[p1] == "1" {
			sum += 1
		}

		if p2 < n && op2[p2] == "1" {
			sum += 1
		}

		switch sum {
		case 0:
			ans = append(ans, "0")
			carry = 0
		case 1:
			ans = append(ans, "1")
			carry = 0
		case 2:
			ans = append(ans, "0")
			carry = 1
		case 3:
			ans = append(ans, "1")
			carry = 1
		}

		p1++
		p2++
	}

	if carry == 1 {
		ans = append(ans, "1")
	}

	reverseStringArray(&ans)

	return strings.Join(ans, "")
}

func reverseStringArray(s *[]string) {
	for i := 0; i < len(*s)/2; i++ {
		(*s)[i], (*s)[len(*s)-i-1] = (*s)[len(*s)-i-1], (*s)[i]
	}
}
