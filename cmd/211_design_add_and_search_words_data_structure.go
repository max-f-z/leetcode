package main

type TN struct {
	children []*TN
	endHere  bool
}

type WordDictionary struct {
	root *TN
}

func ConstructorWD() WordDictionary {
	return WordDictionary{
		&TN{
			children: make([]*TN, 26),
			endHere:  false,
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	ptr := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if ptr.children[ch-'a'] == nil {
			ptr.children[ch-'a'] = &TN{
				children: make([]*TN, 26),
				endHere:  false,
			}
		}
		ptr = ptr.children[ch-'a']
	}
	ptr.endHere = true
}

func searchWord(i int, word string, root *TN) bool {
	if i == len(word) {
		return root.endHere
	}
	ch := word[i]
	if ch != '.' {
		if root.children[ch-'a'] != nil {
			return searchWord(i+1, word, root.children[ch-'a'])
		} else {
			return false
		}
	}
	for j := 0; j < 26; j++ {
		if root.children[j] != nil && searchWord(i+1, word, root.children[j]) {
			return true
		}
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	return searchWord(0, word, this.root)
}
