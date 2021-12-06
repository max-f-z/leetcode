package main

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
