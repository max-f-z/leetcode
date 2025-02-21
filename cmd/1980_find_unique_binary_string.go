package main

func findDifferentBinaryString(nums []string) string {
	n := len(nums)

	ref := map[string]bool{}

	for _, str := range nums {
		ref[str] = true
	}

	return findDifferentBinaryStringHelper("", n, ref)
}

func findDifferentBinaryStringHelper(pfx string, n int, ref map[string]bool) string {
	if len(pfx) == n {
		if !ref[pfx] {
			return pfx
		}
		return ""
	}

	ans := findDifferentBinaryStringHelper(pfx+"0", n, ref)

	if len(ans) == n {
		return ans
	}

	ans = findDifferentBinaryStringHelper(pfx+"1", n, ref)

	return ans
}
