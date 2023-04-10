package main

//lint:ignore U1000 unused
func isOneEditDistance(s string, t string) bool {
	if s == t {
		return false
	}

	switch len(s) - len(t) {
	case 0:
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				return s[i+1:] == t[i+1:]
			}
		}
	case 1:
		for i := 0; i < len(s)-1; i++ {
			if s[i] != t[i] {
				return s[i+1:] == t[i:]
			}
		}
	case -1:
		for i := 0; i < len(t)-1; i++ {
			if s[i] != t[i] {
				return s[i:] == t[i+1:]
			}
		}
	default:
		return false
	}

	return true
}
