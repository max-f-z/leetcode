package main

import "sort"

func minDeletions(s string) int {
	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	arr := make([]int, len(freq))
	idx := 0
	for _, v := range freq {
		arr[idx] = v
		idx++
	}

	sort.Ints(arr)

	seen := map[int]int{}

	for idx := 0; idx < len(arr); idx++ {
		seen[arr[idx]]++
	}

	ans := 0

	for idx = len(arr) - 1; idx >= 0; idx-- {
		if seen[arr[idx]] == 1 || seen[arr[idx]] == 0 {
			continue
		}

		for j := arr[idx] - 1; j >= 0; j-- {
			if seen[j] == 0 || j == 0 {
				ans += arr[idx] - j
				seen[arr[idx]]--
				seen[j] = 1
				break
			}
		}

	}

	return ans
}
