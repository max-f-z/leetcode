package main

//lint:ignore U1000 unused
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	res := -1

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return res
}
