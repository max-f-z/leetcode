package main

import "strconv"

func generateAbbreviations(word string) []string {
	res := []string{}
	generateAbbreviationsHelper(word, 0, "", &res)
	return res
}
func generateAbbreviationsHelper(word string, abbrCount int, temp string, res *[]string) {
	if len(word) == 0 {
		if abbrCount != 0 {
			*res = append(*res, temp+strconv.Itoa(abbrCount))
		} else {
			*res = append(*res, temp)
		}
		return
	}

	generateAbbreviationsHelper(word[1:], abbrCount+1, temp, res)
	if abbrCount == 0 {
		generateAbbreviationsHelper(word[1:], 0, temp+word[:1], res)
	} else {
		generateAbbreviationsHelper(word[1:], 0, temp+strconv.Itoa(abbrCount)+word[:1], res)
	}
}
