package main

func ReverseStringUsingRecursion(s string) string {
	return reverseStringRecursion(s, "")
}

func reverseStringRecursion(s string, prefix string) string {
	if s == "" {
		return prefix
	}

	return reverseStringRecursion(s[:len(s)-1], prefix+string(s[len(s)-1]))
}

func CheckPalindromeRecursion(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return CheckPalindromeRecursion(s[1 : len(s)-1])
}
