package main

import (
	"fmt"
	"strings"
)

func originalDigits(s string) string {

	// zero  count['z']
	// one   count['o'] - count['z'] - count['w'] - count['u']
	// two   count['w']
	// three count['r'] - count['z'] - count['u']
	// four  count['u']
	// five  count['f'] - count['u']
	// six   count['x']
	// seven count['s'] - count['x']
	// eight count['g']
	// nine  (count['n'] - ones - sevens ) / 2

	hash := map[byte]int{}
	for i := range s {
		hash[s[i]]++
	}

	results := make([]int, 10)
	results[0] = hash['z']
	results[2] = hash['w']
	results[4] = hash['u']
	results[5] = hash['f'] - results[4]
	results[6] = hash['x']
	results[8] = hash['g']
	results[1] = hash['o'] - results[0] - results[2] - results[4]
	results[3] = hash['r'] - results[0] - results[4]
	results[7] = hash['s'] - results[6]
	results[9] = hash['i'] - results[5] - results[6] - results[8]

	sb := strings.Builder{}

	for i := 0; i < 10; i++ {
		for j := 0; j < results[i]; j++ {
			sb.WriteString(fmt.Sprint(i))
		}
	}

	return sb.String()
}
