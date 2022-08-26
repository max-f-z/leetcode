package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func reorderedPowerOf2(n int) bool {
	dict := make(map[string]bool)

	for i := 0; i < 30; i++ {
		num := int(math.Pow(2, float64(i)))
		arr := strings.Split(strconv.Itoa(num), "")
		sort.Strings(arr)

		converted := strings.Join(arr, "")

		dict[converted] = true
	}

	queryArr := strings.Split(strconv.Itoa(n), "")
	sort.Strings(queryArr)
	return dict[strings.Join(queryArr, "")]
}
