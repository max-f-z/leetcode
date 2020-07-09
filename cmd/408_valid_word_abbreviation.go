package main

import "strconv"

func validWordAbbreviation(word string, abbr string) bool {
	idx2 := 0
	idx1 := 0

	for idx1 = 0; idx1 < len(abbr); idx1++ {
		if idx2 >= len(word) {
			return false
		}
		if abbr[idx1] >= 49 && abbr[idx1] <= 57 {
			j := 0
			for j = idx1 + 1; j < len(abbr); j++ {
				if abbr[j] >= 47 && abbr[j] <= 57 {
					continue
				}
				break
			}
			num, _ := strconv.Atoi(abbr[idx1:j])
			idx2 += num
			idx1 = j - 1
		} else {
			if word[idx2] == abbr[idx1] {
				idx2++
				continue
			} else {
				return false
			}
		}
	}

	if idx1 == len(abbr) && idx2 == len(word) {
		return true
	}
	return false
}
