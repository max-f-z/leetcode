package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestReverseStringRecursion(t *testing.T) {
	str := "Hello World!"
	ans := reverseStringRecursion(str, "")

	assert.Check(t, ans == reverseString(str))

}

func TestCheckPalindromeRecursion(t *testing.T) {
	str := "Hel  leH"
	ans := CheckPalindromeRecursion(str)

	assert.Check(t, ans)
}
