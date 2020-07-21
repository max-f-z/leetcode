package main

import (
	"fmt"
	"math"
)

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

	fmt.Println(bal)

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

		fmt.Println(bal)

		tmp := minTransfersDFS(start+1, bal)
		if tmp+1 < res {
			res = tmp + 1
		}

		bal[i] -= bal[start]
	}

	return res
}
