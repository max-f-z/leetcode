package main

import (
	"fmt"
	"strconv"
)

func findContestMatch(n int) string {
	strs := []string{}

	for i := 1; i < n+1; i++ {
		strs = append(strs, strconv.Itoa(i))
	}

	return findContestMatchHelper(strs)
}

func findContestMatchHelper(strs []string) string {
	resultArr := []string{}

	for i := 0; i < len(strs)/2; i++ {
		resultArr = append(resultArr, fmt.Sprintf("(%s,%s)", strs[i], strs[len(strs)-1-i]))
	}

	if len(resultArr) != 1 {
		resultArr[0] = findContestMatchHelper(resultArr)
	}

	return resultArr[0]
}
