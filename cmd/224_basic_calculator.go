package main

import (
	"container/list"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
func calculate(s string) int {
	result := 0
	s = strings.ReplaceAll(s, " ", "")

	stack := list.New().Init()

	prefix := strings.Builder{}
	for idx := range s {
		if s[idx] == '(' {
			stack.PushBack(prefix.String())
			prefix.Reset()

		} else if s[idx] == ')' {
			val := calculateHelper(prefix.String())

			prefix.Reset()
			back := stack.Back()
			stack.Remove(back)

			backStr := back.Value.(string)
			if val < 0 && len(backStr) > 0 && backStr[len(backStr)-1] == '-' {
				backStr = backStr[:len(backStr)-1] + "+" + strconv.Itoa(val*-1)
			} else if val < 0 && len(backStr) > 0 && backStr[len(backStr)-1] == '+' {
				backStr = backStr[:len(backStr)-1] + "-" + strconv.Itoa(val*-1)
			} else {
				backStr += strconv.Itoa(val)
			}

			prefix.WriteString(backStr)

		} else {
			prefix.WriteByte(s[idx])
		}
	}

	result = calculateHelper(prefix.String())

	return result
}

func calculateHelper(s string) int {
	if s == "" {
		return 0
	}

	prefix := strings.Builder{}
	sum := 0
	operator := 1
	for idx := range s {
		if s[idx] == '+' {
			val, _ := strconv.Atoi(prefix.String())
			sum += val * operator
			operator = 1
			prefix.Reset()
		} else if s[idx] == '-' {
			val, _ := strconv.Atoi(prefix.String())
			sum += val * operator
			operator = -1
			prefix.Reset()
		} else {
			prefix.WriteByte(s[idx])
		}
	}

	if len(prefix.String()) != 0 {
		val, _ := strconv.Atoi(prefix.String())
		sum += val * operator
		prefix.Reset()
	}

	return sum
}
