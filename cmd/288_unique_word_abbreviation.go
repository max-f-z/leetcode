package main

import "strconv"

type ValidWordAbbr struct {
	words []string
	abbrs []string
}

func Constructor288(dictionary []string) ValidWordAbbr {
	dict := ValidWordAbbr{}
	dict.words = dictionary
	dict.abbrs = []string{}
	for _, v := range dictionary {
		tmp := ""
		if len(v) < 3 {
			tmp = v
		} else {
			tmp = string(v[0]) + strconv.Itoa(len(v)-2) + string(v[len(v)-1])
		}
		dict.abbrs = append(dict.abbrs, tmp)
	}

	return dict
}

//lint:ignore ST1006 this
func (this *ValidWordAbbr) IsUnique(word string) bool {
	found := false

	abbr := ""
	if len(word) < 3 {
		abbr = word
	} else {
		abbr = string(word[0]) + strconv.Itoa(len(word)-2) + string(word[len(word)-1])
	}

	for k, v := range this.abbrs {
		if v == abbr {
			found = true
			if this.words[k] != word {
				return false
			}
		}
	}

	if !found {
		return true
	}

	return true
}
