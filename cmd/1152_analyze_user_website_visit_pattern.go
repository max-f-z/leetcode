package main

import (
	"sort"
	"strings"
)

type tuple struct {
	username  string
	timestamp int
	website   string
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {

	tuples := make([]tuple, len(username))

	for i := 0; i < len(username); i++ {
		tuples[i] = tuple{
			username:  username[i],
			timestamp: timestamp[i],
			website:   website[i],
		}
	}

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].timestamp < tuples[j].timestamp
	})

	vals := map[string][]string{}

	for i := 0; i < len(website); i++ {
		if vals[tuples[i].username] == nil {
			vals[tuples[i].username] = []string{}
		}

		vals[tuples[i].username] = append(vals[tuples[i].username], tuples[i].website)
	}

	cnt := map[string]int{}

	for _, v := range vals {
		cntU := map[string]int{}
		for i := 0; i < len(v)-2; i++ {
			for j := i + 1; j < len(v)-1; j++ {
				for k := j + 1; k < len(v); k++ {
					sb := strings.Builder{}
					sb.WriteString(v[i])
					sb.WriteByte(' ')
					sb.WriteString(v[j])
					sb.WriteByte(' ')
					sb.WriteString(v[k])

					cntU[sb.String()]++
				}
			}
		}

		for k := range cntU {
			cnt[k]++
		}
	}

	max := 0
	resStr := ""
	var res []string
	for k, v := range cnt {
		if v > max {
			resStr = k
			res = strings.Split(k, " ")
			max = v
		}

		if v == max {
			tmp := k
			if tmp < resStr {
				resStr = tmp
				res = strings.Split(k, " ")
			}
		}
	}

	return res
}
