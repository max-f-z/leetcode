package main

import (
	"strings"
)

// https://www.youtube.com/watch?v=ZU7fiX0WCCY
//
//lint:ignore U1000 unused
func alienOrder(words []string) string {
	parent := make(map[byte]byte, 0)

	for i := 0; i < len(words); i++ {
		for j := range words[i] {
			parent[words[i][j]] = 0
		}
	}

	for i := 0; i < len(words)-1; i++ {
		str1, str2 := words[i], words[i+1]

		if len(str1) > len(str2) && strings.HasPrefix(str1, str2) {
			return ""
		}

		for j := 0; j < len(str1); j++ {
			if str1[j] != str2[j] {
				if parent[str1[j]] == str2[j] {
					return ""
				}
				parent[str2[j]] = str1[j]

				break
			}
		}
	}

	res := ""

	for len(parent) > 0 {
		var remove byte
		found := false
		for k := range parent {
			if parent[k] == 0 {
				delete(parent, k)
				remove = k
				res += string(k)
				found = true
				break
			}
		}

		if !found {
			return ""
		}

		for k := range parent {
			if parent[k] == remove {
				parent[k] = 0
			}
		}
	}

	return res
}

//lint:ignore U1000 unused
func alienOrder2(words []string) string {
	parent := make(map[byte][]byte, 0)

	for i := 0; i < len(words); i++ {
		for j := range words[i] {
			parent[words[i][j]] = []byte{}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		str1, str2 := words[i], words[i+1]

		if len(str1) > len(str2) && strings.HasPrefix(str1, str2) {
			return ""
		}

		for j := 0; j < len(str1); j++ {
			if str1[j] != str2[j] {
				parent[str2[j]] = append(parent[str2[j]], str1[j])
				break
			}
		}
	}

	res := ""

	for len(parent) > 0 {
		var remove byte
		found := false
		for k := range parent {
			if len(parent[k]) == 0 {
				delete(parent, k)
				remove = k
				res += string(k)
				found = true
				break
			}
		}

		if !found {
			return ""
		}

		for k := range parent {
			for j := range parent[k] {
				if parent[k][j] == remove {
					parent[k] = removeFromSlice(parent[k], remove)
				}
			}
		}
	}

	return res
}

func removeFromSlice(arr []byte, el byte) []byte {
	if arr[0] == el {
		return arr[0:]
	}
	if arr[len(arr)-1] == el {
		return arr[:len(arr)-1]
	}
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] == el {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return []byte{}
}
