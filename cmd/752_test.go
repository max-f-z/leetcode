package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test752(t *testing.T) {
	deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target := "8888"

	ans := openLock(deadends, target)

	assert.Check(t, ans == 6)
}
