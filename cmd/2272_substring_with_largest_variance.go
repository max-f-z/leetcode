package main

//lint:ignore U1000 unused
func largestVariance(s string) int {
	max := 0
	var a, b byte
	for a = 'a'; a <= 'z'; a++ {
		for b = 'a'; b <= 'z'; b++ {
			val := solveOne(a, b, s)
			if val > max {
				max = val
			}
		}
	}
	return max
}

func solveOne(a, b byte, s string) int {
	if a == b {
		return 0
	}

	max := 0
	hasB := false
	bFlag := false

	local := 0

	for i := 0; i < len(s); i++ {
		if s[i] == a {
			local++
		} else if s[i] == b {
			if bFlag && local >= 0 {
				bFlag = false
			} else {
				local--
			}

			hasB = true

			if local < 0 {
				local = -1
				bFlag = true
			}
		}

		if local > max && hasB {
			max = local
		}
	}
	return max
}
