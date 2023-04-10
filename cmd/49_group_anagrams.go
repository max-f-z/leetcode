package main

//lint:ignore U1000 unused
func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}

	groups := []map[byte]int{}

	for _, str := range strs {
		group := make(map[byte]int)
		for i := range str {
			group[str[i]]++
		}

		find := -1
		for i := range groups {
			if len(groups[i]) == len(group) {
				flag := true
				for k, v := range group {
					if groups[i][k] != v {
						flag = false
						break
					}
				}

				if flag {
					find = i
					break
				}
			}
		}

		if find != -1 {
			ans[find] = append(ans[find], str)
		} else {
			groups = append(groups, group)
			ans = append(ans, []string{str})
		}
	}

	return ans
}

//lint:ignore U1000 unused
func groupAnagramsRefactor(strs []string) [][]string {
	groups := map[[26]int][]string{}

	for _, str := range strs {
		cnt := [26]int{}

		for i := 0; i < len(str); i++ {
			cnt[str[i]-'a'] += 1
		}

		if groups[cnt] == nil {
			groups[cnt] = []string{str}
		} else {
			groups[cnt] = append(groups[cnt], str)
		}
	}

	ans := make([][]string, len(groups))
	idx := 0
	for _, v := range groups {
		ans[idx] = v
		idx++
	}

	return ans
}
