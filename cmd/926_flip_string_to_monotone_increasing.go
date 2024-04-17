package main

func minFlipsMonoIncr(s string) int {
	// ans keep the final ans for each starting[:i] substring
	ans := 0
	// numOf1s keep the num of 1s before i
	numOf1s := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			// if s[i] is 0,
			// final = min(final + 1, numOf1s)
			// final + 1, flip current
			// numOf1s, flip all previous 1s
			ans += 1
			if ans > numOf1s {
				ans = numOf1s
			}
		} else {
			// if s[i] is 1, no need to flip final = final previous
			numOf1s++
		}
	}

	return ans
}
