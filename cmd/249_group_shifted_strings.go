package main

func groupStrings(strings []string) [][]string {
	res := [][]string{}

	for _, v := range strings {
		found := false
		for i := range res {
			if len(res[i]) > 0 && sameShiftedString(res[i][0], v) {
				res[i] = append(res[i], v)
				found = true
			}
		}
		if !found {
			res = append(res, []string{v})
		}
	}

	return res
}

func sameShiftedString(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 1 {
		return true
	}

	for i := 0; i < len(s1)-1; i++ {
		tmp1 := int(s1[i]) - int(s1[i+1])
		tmp2 := int(s2[i]) - tmp1
		if tmp2 < 97 {
			tmp2 += 26
		}
		if tmp2 > 122 {
			tmp2 -= 26
		}
		if tmp2 != int(s2[i+1]) {
			return false
		}
	}

	return true
}
