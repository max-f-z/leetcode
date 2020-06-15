package main

func romanToInt(s string) int {
	ref := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	res := 0
	flag := false
	for i := 0; i < len(s); i++ {
		if i < (len(s) - 1) {
			if ref[string(s[i])] < ref[string(s[i+1])] && !flag {
				res -= ref[string(s[i])]
				flag = true
			} else {
				res += ref[string(s[i])]
				flag = false
			}
		} else {
			res += ref[string(s[i])]
			flag = false
		}
	}

	return res
}
