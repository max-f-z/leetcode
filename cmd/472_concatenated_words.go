package main

func findAllConcatenatedWordsInADict(words []string) []string {
	cols := map[string]int{}

	for i := 0; i < len(words); i++ {
		cols[words[i]] = 1
	}

	res := []string{}

	for i := 0; i < len(words); i++ {
		delete(cols, words[i])

		val := findAllConcatenatedWordsInADictHelper(cols, words[i])

		if val > 1 {
			res = append(res, words[i])
		} else if val == 0 {
			cols[words[i]] = 1
		}
	}

	return res
}

func findAllConcatenatedWordsInADictHelper(set map[string]int, word string) int {
	if v, ok := set[word]; ok {
		return v
	}

	for i := 1; i < len(word); i++ {
		if set[word[0:i]] > 0 {
			val := findAllConcatenatedWordsInADictHelper(set, word[i:])
			if val > 0 {
				set[word] = val + set[word[0:i]]
				return set[word]
			}
		}
	}
	// 关键减少重复运算
	set[word] = 0
	return 0
}
