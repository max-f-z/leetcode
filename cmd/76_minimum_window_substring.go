package main

func minWindow(s string, t string) string {
	start, end := 0, -1

	sl, tl := len(s), len(t)

	count := map[byte]int{}

	for i := 0; i < tl; i++ {
		count[t[i]]++
	}

	j := 0
	need := tl
	for i := 0; i < sl; i++ {
		if count[s[i]] > 0 {
			need--
		}

		count[s[i]]--

		if need == 0 {
			for j < i && count[s[j]] < 0 {
				count[s[j]]++
				j++
			}

			if end == -1 || end-start > i-j {
				start, end = j, i
			}

			need++
			count[s[j]]++
			j++
		}
	}

	if end < start {
		return ""
	}

	return s[start : end+1]
}
