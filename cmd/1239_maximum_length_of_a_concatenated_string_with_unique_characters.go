package main

func checkUnique(s string) bool {
	m := map[byte]int{}

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] > 1 {
			return false
		}
	}
	return true
}

func maxLength(arr []string) int {
	return maxLengthHelper(0, arr, "")
}

func maxLengthHelper(start int, arr []string, prev string) int {
	if len(prev) > 26 {
		return 0
	}

	if start > len(arr)-1 {
		if checkUnique(prev) {
			return len(prev)
		}
	}

	noCur := maxLengthHelper(start+1, arr, prev)
	haveCur := 0
	if checkUnique(prev + arr[start]) {
		haveCur = maxLengthHelper(start+1, arr, prev+arr[start])
	}

	if noCur > haveCur {
		return noCur
	}

	return haveCur
}
