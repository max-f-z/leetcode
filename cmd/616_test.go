package main

import "testing"

func Test616(t *testing.T) {
	s := addBoldTag("aaabbcc", []string{"aaa", "aab", "bc"})

	t.Log(s)

	s = addBoldTag("abcxyz123", []string{"abc", "123", "bc"})

	t.Log(s)
}
