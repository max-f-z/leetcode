package main

//lint:ignore U1000 unused
func wordSquares(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}

	length := len(words[0])

	res := [][]string{}

	trie := ConstructTrie(words)

	for i := 0; i < len(words); i++ {
		searchWordSquare(trie, length, 1, []string{words[i]}, &res)
	}

	return res
}

func searchWordSquare(node *TrieNode, length int, idx int, tmp []string, res *[][]string) {
	if len(tmp) == length {
		cp := make([]string, length)
		copy(cp, tmp)
		(*res) = append(*res, cp)
		return
	}

	prefix := ""
	for i := 0; i < len(tmp); i++ {
		prefix += string((tmp)[i][idx])
	}

	available := node.startsWith(prefix)

	for i := 0; i < len(available); i++ {
		tmpAva := append(tmp, available[i])
		searchWordSquare(node, length, idx+1, tmpAva, res)
	}

}

type TrieNode struct {
	children []*TrieNode
	isWord   bool
}

func ConstructTrie(strs []string) *TrieNode {
	root := &TrieNode{}
	root.children = make([]*TrieNode, 26)

	for _, str := range strs {
		cur := root
		for _, v := range str {
			if cur.children[int(v)-97] == nil {
				cur.children[int(v)-97] = &TrieNode{children: make([]*TrieNode, 26)}
			}

			cur = cur.children[int(v)-97]
		}
		cur.isWord = true
	}

	return root
}

//lint:ignore U1000 unused
func (root *TrieNode) search(word string) bool {
	cur := root

	for _, v := range word {
		if cur.children[int(v)-97] == nil {
			return false
		}
		cur = cur.children[int(v)-97]
	}

	return cur.isWord
}

func (root *TrieNode) startsWith(word string) []string {
	res := []string{}
	cur := root
	for _, v := range word {
		if cur.children[int(v)-97] == nil {
			return res
		}
		cur = cur.children[int(v)-97]
	}

	if cur.isWord {
		res = append(res, word)
	}

	startsWithHelper(cur, word, &res)

	return res
}

func startsWithHelper(node *TrieNode, prefix string, res *[]string) {
	if node == nil {
		return
	}

	if node.isWord {
		*res = append(*res, prefix)
	}

	for i, v := range node.children {
		if v != nil {
			startsWithHelper(v, prefix+string(byte(i+97)), res)
		}
	}
}
