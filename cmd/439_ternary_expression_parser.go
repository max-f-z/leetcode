package main

import "strings"

//lint:ignore U1000 unused
func parseTernary(expression string) string {
	return parseTernaryHelper(expression)
}

func parseTernaryHelper(expression string) string {
	idx := strings.LastIndexByte(expression, '?')

	if idx == -1 {
		return expression
	}

	nextColonIdx := -1
	for i := idx + 1; i < len(expression); i++ {
		if expression[i] == ':' {
			nextColonIdx = i
			break
		}
	}

	boundry := -1
	for boundry = nextColonIdx + 1; boundry < len(expression); boundry++ {
		if expression[boundry] == ':' {
			break
		}
	}

	if expression[idx-1] == 'T' {
		if boundry == len(expression)-1 {
			newExpression := expression[:idx-1] + expression[idx+1:nextColonIdx]
			return parseTernaryHelper(newExpression)
		}
		newExpression := expression[:idx-1] + expression[idx+1:nextColonIdx] + expression[boundry:]
		return parseTernaryHelper(newExpression)

	}

	newExpression := expression[:idx-1] + expression[nextColonIdx+1:]
	return parseTernaryHelper(newExpression)
}
