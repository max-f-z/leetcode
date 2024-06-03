package main

func appendCharacters(s string, t string) int {
	p1, p2 := 0, 0

	for p1 = 0; p1 < len(s); {
		if p2 == len(t) {
			break
		}

		if s[p1] == t[p2] {
			p2++
		}
		p1++
	}

	return len(t) - p2
}
