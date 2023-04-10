package main

//lint:ignore U1000 unused
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	index, cnt := 0, 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[index] {
			cnt++
			index++
		}

		if cnt == len(s) {
			return true
		}
	}

	return false
}
