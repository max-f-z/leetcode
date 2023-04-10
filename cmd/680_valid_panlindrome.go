package main

//lint:ignore U1000 unused
func validPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] == s[len(s)-1-i] {
			continue
		}

		return isRangeValid(s, i+1, len(s)-1-i) || isRangeValid(s, i, len(s)-2-i)
	}

	return true
}

func isRangeValid(s string, i, j int) bool {
	for k := i; k <= i+(j-i)/2; k++ {
		if s[k] != s[j-k+i] {
			return false
		}
	}
	return true
}
