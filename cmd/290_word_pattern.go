package main

import "strings"

//lint:ignore U1000 unused
func wordPattern(pattern string, str string) bool {
	keys := map[string]string{}
	vals := map[string]string{}
	strs := strings.Split(str, " ")

	if len(strs) != len(pattern) {
		return false
	}

	for k, v := range strs {
		if keys[string(pattern[k])] == "" && vals[v] == "" {
			keys[string(pattern[k])] = v
			vals[v] = string(pattern[k])
			continue
		}

		if keys[string(pattern[k])] != "" && keys[string(pattern[k])] != v {
			return false
		}

		if vals[v] != "" && vals[v] != string(pattern[k]) {
			return false
		}
	}

	return true
}
