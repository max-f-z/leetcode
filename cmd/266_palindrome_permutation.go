package main

func canPermutePalindrome(s string) bool {
	dict := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	even := 0
	odd := 0

	for _, v := range dict {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if len(s)%2 == 0 {
		if odd == 0 {
			return true
		}
	} else {
		if odd == 1 {
			return true
		}
	}

	return false
}
