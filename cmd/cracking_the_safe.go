package main

import (
	"fmt"
	"math"
)

func crackSafe(n int, k int) string {
	res := ""

	for i := n; i > 0; i-- {
		res += "0"
	}

	set := map[string]int{}
	set[res] = 1

	for i := float64(0); i < math.Pow(float64(k), float64(n)); i++ {
		tmp := string(res[len(res)-n+1:])
		for j := k - 1; j >= 0; j-- {
			if set[fmt.Sprintf("%s%d", tmp, j)] == 0 {
				set[fmt.Sprintf("%s%d", tmp, j)] = 1
				res = fmt.Sprintf("%s%d", res, j)
				break
			}
		}
	}

	return res
}
