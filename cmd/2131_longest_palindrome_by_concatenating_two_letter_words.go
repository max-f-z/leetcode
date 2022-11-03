package main

func longestPalindrome2131(words []string) int {
	counts := map[string]int{}

	for _, word := range words {
		counts[word]++
	}

	ans := 0
	central := false
	for k := range counts {
		if k[0] == k[1] {
			if counts[k]%2 == 0 {
				ans += counts[k]
			} else {
				ans += (counts[k] - 1)
				central = true
			}
		} else if k[0] < k[1] && counts[string(k[1])+string(k[0])] > 0 {
			min := counts[k]
			if counts[k] > counts[string(k[1])+string(k[0])] {
				min = counts[string(k[1])+string(k[0])]
			}
			ans += 2 * min
		}
	}
	if central {
		ans++
	}

	return ans * 2
}
