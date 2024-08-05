package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2053(t *testing.T) {
	arr := []string{"aaa", "aa", "a"}
	ans := kthDistinct(arr, 1)
	assert.Check(t, ans == "aaa")
}
