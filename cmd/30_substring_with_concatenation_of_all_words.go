package main

//lint:ignore U1000 unused
func findSubstring(s string, words []string) []int {
	res := []int{}

	if s == "" {
		return []int{}
	}

	if len(words) == 0 {
		return []int{}
	}

	for i := 0; i <= len(s)-(len(words[0])*len(words)); i++ {
		if helper(s[i:], words) {
			res = append(res, i)
		}
	}
	return res
}

func helper(s string, words []string) bool {
	if len(words) == 0 {
		return true
	}
	step := len(words[0])

	copy := []string{}
	for _, v := range words {
		copy = append(copy, v)
	}

	for i, v := range words {
		if v == s[:step] {
			if i == 0 {
				copy = copy[1:]
			} else if i == len(words)-1 {
				copy = copy[:len(words)-1]
			} else {
				copy = append(copy[:i], copy[i+1:]...)
			}

			return helper(s[step:], copy)
		}
	}

	return false
}
