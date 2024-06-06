package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test846(t *testing.T) {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	ans := isNStraightHand(hand, 3)
	assert.Check(t, ans == true)

	hand = []int{1, 2, 3, 4, 5, 6}
	ans = isNStraightHand(hand, 2)
	assert.Check(t, ans == true)

	hand = []int{1, 2, 3}
	ans = isNStraightHand(hand, 1)
	assert.Check(t, ans == true)
}
