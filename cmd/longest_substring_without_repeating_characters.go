package main

//lint:ignore U1000 unused
func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		tmp := 0
		set := map[string]int{}
		for j := i; j < len(s); j++ {
			if set[string(s[j])] == 0 {
				tmp++
				set[string(s[j])] = 1
				continue
			}
			break
		}
		if tmp > max {
			max = tmp
		}
	}

	return max
}
