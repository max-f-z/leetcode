package main

func commonChars(words []string) []string {
	ans := map[string]int{}

	for i := 0; i < len(words[0]); i++ {
		ans[string(words[0][i])]++
	}

	for i := 1; i < len(words); i++ {
		cnt := map[string]int{}
		for j := 0; j < len(words[i]); j++ {
			cnt[string(words[i][j])]++
		}

		for str := range ans {
			if cnt[str] == 0 {
				delete(ans, str)
			} else if ans[str] > cnt[str] {
				ans[str] = cnt[str]
			}
		}
	}
	keys := []string{}

	for str := range ans {
		for i := 0; i < ans[str]; i++ {
			keys = append(keys, str)
		}
	}
	return keys
}
