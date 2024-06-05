package main

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
)

func Test1002(t *testing.T) {
	words := []string{"bella", "label", "roller"}

	ans := commonChars(words)
	sort.Strings(ans)

	assert.DeepEqual(t, ans, []string{"e", "l", "l"})
}
