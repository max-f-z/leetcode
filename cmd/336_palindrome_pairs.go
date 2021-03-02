package main

func palindromePairs(words []string) [][]int {
	pairs := map[string]int{}

	res := [][]int{}

	for i := range words {
		pairs[reverseStr(words[i])] = i + 1
	}

	for i := range words {
		for j := 0; j < len(words[i]); j++ {
			left, right := words[i][0:j], words[i][j:]

			if isPalindromeStr(left) && pairs[right] != 0 && pairs[right] != i+1 {
				res = append(res, []int{pairs[right] - 1, i})
			}

			if isPalindromeStr(right) && pairs[left] != 0 && pairs[left] != i+1 {
				res = append(res, []int{i, pairs[left] - 1})
				if left == "" {
					res = append(res, []int{pairs[left] - 1, i})
				}
			}
		}
	}

	return res
}

func reverseStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindromeStr(str string) bool {
	l := len(str)

	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}

	return true
}
