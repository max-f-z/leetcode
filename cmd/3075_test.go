package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test3075(t *testing.T) {
	happiness := []int{1, 2, 3}
	ans := maximumHappinessSum(happiness, 2)
	assert.Check(t, ans == 4)

	happiness = []int{1, 1, 1, 1}
	ans = maximumHappinessSum(happiness, 2)
	assert.Check(t, ans == 1)
}
