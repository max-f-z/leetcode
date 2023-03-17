package main

import "testing"

func Test208(t *testing.T) {
	trie := ConstructorTrie()
	trie.Insert("apple")

	res := trie.Search("app")
	t.Log(res)

	res = trie.StartsWith("app")
	t.Log(res)

	res = trie.Search("apple")
	t.Log(res)
}
