package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test1190(t *testing.T) {
	s := "(abcd)"
	ans := reverseParentheses(s)
	assert.Check(t, ans == "dcba")

	s = "(u(love)i)"
	ans = reverseParentheses(s)
	assert.Check(t, ans == "iloveu")

	s = "(ed(et(oc))el)"
	ans = reverseParentheses(s)
	assert.Check(t, ans == "leetcode")
}
