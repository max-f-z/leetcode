package main

import (
	"math"
)

//lint:ignore U1000 unused
func minTransfers(transactions [][]int) int {
	tmp := make(map[int]int)

	for _, v := range transactions {
		tmp[v[0]] -= v[2]
		tmp[v[1]] += v[2]
	}

	bal := []int{}
	for _, v := range tmp {
		if v != 0 {
			bal = append(bal, v)
		}
	}
	if len(bal) == 0 {
		return 0
	}

	return minTransfersDFS(0, bal)
}

func minTransfersDFS(start int, bal []int) int {
	for start < len(bal) && bal[start] == 0 {
		start++
	}

	if start == len(bal) {
		return 0
	}

	if len(bal)-start == 2 {
		return 1
	}

	res := math.MaxInt32

	for i := start + 1; i < len(bal); i++ {
		bal[i] += bal[start]

		tmp := minTransfersDFS(start+1, bal)
		if tmp+1 < res {
			res = tmp + 1
		}

		bal[i] -= bal[start]
	}

	return res
}
