package main

import "fmt"

// https://www.youtube.com/watch?v=Vbcl5jmm05E

func generatePalindromes(s string) []string {
	res := make([]string, 0)

	dict := make([]int, 256)

	for i := 0; i < len(s); i++ {
		dict[int(s[i])]++
	}

	even := 0
	odd := 0
	var oddStr byte

	for k := range dict {
		if dict[k]%2 == 0 {
			even++
		} else {
			odd++
			oddStr = byte(k)
		}
	}

	if len(s)%2 == 0 {
		if odd != 0 {
			return res
		}
	} else {
		if odd != 1 {
			return res
		}
		dict[oddStr]--
	}

	oddStrr := ""
	if oddStr != 0 {
		oddStrr = string(oddStr)
	}

	generatePalindromesHelper(&res, dict, oddStrr, len(s))

	return res
}

func generatePalindromesHelper(res *[]string, dict []int, tmp string, n int) {
	if len(tmp) == n {
		*res = append(*res, tmp)
		return
	}

	for k := range dict {
		if dict[k] > 0 {
			dict[k] -= 2

			generatePalindromesHelper(res, dict, fmt.Sprint(k)+tmp+fmt.Sprint(k), n)

			dict[k] += 2
		}
	}
}
