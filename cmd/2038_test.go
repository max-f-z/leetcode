package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2038(t *testing.T) {
	ans := winnerOfGame("ABBBBBBBAAA")
	assert.Check(t, ans == false)

	ans = winnerOfGame("AA")
	assert.Check(t, ans == false)

	ans = winnerOfGame("AAABABB")
	assert.Check(t, ans == true)

}
