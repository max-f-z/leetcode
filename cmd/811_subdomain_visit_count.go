package main

import (
	"fmt"
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
func subdomainVisits(cpdomains []string) []string {
	tmp := map[string]int{}

	for _, v := range cpdomains {
		visit := strings.Split(v, " ")

		num, _ := strconv.Atoi(visit[0])

		domains := strings.Split(visit[1], ".")

		for i := len(domains) - 1; i >= 0; i-- {
			tmp[strings.Join(domains[i:], ".")] += num
		}
	}

	result := []string{}
	for k := range tmp {
		result = append(result, fmt.Sprintf("%d %s", tmp[k], k))
	}

	return result
}
