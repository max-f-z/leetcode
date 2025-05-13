package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3335(t *testing.T) {
	s := "abcyy"
	round := 2
	ans := lengthAfterTransformations(s, round)
	assert.Check(t, ans == 7)

	s = "azbk"
	round = 1
	ans = lengthAfterTransformations(s, round)
	assert.Check(t, ans == 5)

	s = "xly"
	round = 5
	ans = lengthAfterTransformations(s, round)
	assert.Check(t, ans == 5)

}
