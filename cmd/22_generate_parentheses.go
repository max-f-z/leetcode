package main

func generateParenthesis(n int) []string {
	res := []string{}
	backTracking(&res, "", 0, 0, n)
	return res
}

func backTracking(res *[]string, str string, left, right int, n int) {
	if len(str) == 2*n {
		*res = append(*res, str)
	}

	if left > n || right > n {
		return
	}

	if left < right {
		return
	}
	if left < n {
		backTracking(res, str+"(", left+1, right, n)
	}
	if right < n {
		backTracking(res, str+")", left, right+1, n)
	}
}
