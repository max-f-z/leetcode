package main

func lengthOfLongestSubstringTwoDistinct(s string) int {
	max := 0

	box := map[string]int{}
	start := 0
	end := 0
	for end < len(s) {
		if len(box) <= 2 {
			box[string(s[end])]++
			end++
		} else {
			box[string(s[start])]--
			if box[string(s[start])] == 0 {
				delete(box, string(s[start]))
			}
			start++
		}

		if len(box) <= 2 && max < (end-start) {
			max = end - start
		}
	}

	return max
}
