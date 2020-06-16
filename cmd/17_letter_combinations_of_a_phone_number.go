package main

func letterCombinations(digits string) []string {
	dict := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	res := []string{}

	for _, v := range digits {
		tmp := []string{}

		if len(res) == 0 {
			for _, w := range dict[string(v)] {
				tmp = append(tmp, w)
			}
		} else {
			for _, u := range res {
				for _, w := range dict[string(v)] {
					tmp = append(tmp, u+w)
				}
			}
		}

		res = tmp
	}

	return res
}
