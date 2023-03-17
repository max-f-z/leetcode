package main

type trieNode struct {
	nodes map[byte]*trieNode
	isEnd bool
}

type Trie struct {
	root *trieNode
}

func ConstructorTrie() Trie {
	return Trie{
		root: &trieNode{
			nodes: make(map[byte]*trieNode),
			isEnd: false,
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root

	for i := range word {
		if cur.nodes[word[i]] == nil {
			cur.nodes[word[i]] = &trieNode{
				nodes: make(map[byte]*trieNode),
			}
		}

		if i == len(word)-1 {
			cur.nodes[word[i]].isEnd = true
		}
		cur = cur.nodes[word[i]]
	}
}

func (this *Trie) Search(word string) bool {
	cur := this.root

	for i := range word {
		if cur.nodes[word[i]] != nil {
			cur = cur.nodes[word[i]]
			if i == (len(word)-1) && cur.isEnd {
				return true
			}
			continue
		}
		return false
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for i := range prefix {
		if cur.nodes[prefix[i]] != nil {
			cur = cur.nodes[prefix[i]]
			if i == (len(prefix) - 1) {
				return true
			}
			continue
		}
		return false
	}
	return false
}
