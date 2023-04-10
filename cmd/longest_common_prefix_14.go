package main

import "math"

//lint:ignore U1000 unused
func longestCommonPrefix(strs []string) string {
	min := math.MaxInt32
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}

	res := ""

	if len(strs) == 0 {
		return res
	}

	for i := 0; i < min; i++ {

		tmp := string(strs[0][len(res)])

		tmpRes := true
		for j := 0; j < len(strs); j++ {

			tmpRes = tmpRes && (tmp == string(strs[j][len(res)]))
		}

		if tmpRes {
			res += tmp
			continue
		}
		break
	}
	return res
}
