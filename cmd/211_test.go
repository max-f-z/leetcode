package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

// ["WordDictionary","addWord","addWord","search","search","search","search","search","search","search","search"]
// [[],["a"],["ab"],["a"],["a."],["ab"],[".a"],[".b"],["ab."],["."],[".."]]
// [null,null,null,true,true,true,false,true,false,true,true]

func Test211(t *testing.T) {
	wd := ConstructorWD()

	// wd.AddWord("bad")
	// wd.AddWord("dad")
	// wd.AddWord("mad")

	// ans := wd.Search("pad")
	// assert.Check(t, ans == false)
	// ans = wd.Search("bad")
	// assert.Check(t, ans == true)
	// ans = wd.Search(".ad")
	// assert.Check(t, ans == true)
	// ans = wd.Search("b..")
	// assert.Check(t, ans == true)

	wd.AddWord("a")
	wd.AddWord("ab")
	ans := wd.Search("a")
	assert.Check(t, ans == true)
	ans = wd.Search("a.")
	assert.Check(t, ans == true)
	ans = wd.Search("ab")
	assert.Check(t, ans == true)
	ans = wd.Search(".a")
	assert.Check(t, ans == false)
	ans = wd.Search(".b")
	assert.Check(t, ans == true)
	ans = wd.Search("ab.")
	assert.Check(t, ans == false)
	ans = wd.Search(".")
	assert.Check(t, ans == true)
	ans = wd.Search("..")
	assert.Check(t, ans == true)
}
