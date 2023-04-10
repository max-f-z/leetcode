package main

//lint:ignore U1000 unused
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	max := 0

	box := map[string]int{}

	start := 0
	end := 0
	for end < len(s) {
		if len(box) <= k {
			box[string(s[end])]++
			end++
		} else {
			box[string(s[start])]--
			if box[string(s[start])] == 0 {
				delete(box, string(s[start]))
			}
			start++
		}

		if len(box) <= k && max < (end-start) {
			max = end - start
		}
	}
	return max
}
