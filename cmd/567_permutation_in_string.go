package main

func checkInclusion(s1 string, s2 string) bool {
	cnt := map[byte]int{}

	if len(s2) < len(s1) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		cnt[s1[i]]++
	}

	local := map[byte]int{}

	for i := 0; i < len(s1); i++ {
		local[s2[i]]++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		if i != 0 {
			local[s2[i-1]]--
			local[s2[i+len(s1)-1]]++
			if local[s2[i-1]] == 0 {
				delete(local, s2[i-1])
			}
		}
		if checkMatch(local, cnt) {
			return true
		}
	}

	return false
}

func checkMatch(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}
	return true
}
