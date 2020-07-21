package main

import "strings"

func findPermutation(s string) []int {
	res := make([]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		res[i] = i + 1
	}
	if strings.IndexByte(s, 'D') == -1 {
		return res
	}

	if strings.IndexByte(s, 'I') == -1 {
		for i := 0; i < len(s)+1; i++ {
			res[i] = len(s) + 1 - i
		}
		return res
	}

	s += "I"

	tmp := []int{}
	idx := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			tmp = append(tmp, i+1)
			continue
		}

		res[idx] = i + 1
		idx++

		for j := 0; j < len(tmp); j++ {
			res[idx] = tmp[len(tmp)-1-j]
			idx++
		}

		tmp = []int{}
	}

	return res
}

/*
## APPROACH : STACK ##
## LOGIC ##
## 1. When we get D, we push and when we get I, we calculate previous maximum value and calculate the next number i.e prev + 1 and add to current I position and as only D's are in the stack and they have lower precedence than I, we pop all incrementing the prev value updated
## 2. To simplify the things, I have appended I at the end. ( if there are any D's left in the stack, this will take care and also we are missing out the last number, i.e len(s), so include that aswell we need I at the end)

## TIME COMPLEXITY : O(N) ##
## SPACE COMPLEXITY : O(N) ##

## EXAMPLE : "DDIIDDID"	##
## STACK TRACE ##
# D [0, 0, 0, 0, 0, 0, 0, 0, 0] [('D', 0)]
# D [0, 0, 0, 0, 0, 0, 0, 0, 0] [('D', 0), ('D', 1)]
# I [3, 2, 1, 0, 0, 0, 0, 0, 0] []
# I [3, 2, 1, 4, 0, 0, 0, 0, 0] []
# D [3, 2, 1, 4, 0, 0, 0, 0, 0] [('D', 4)]
# D [3, 2, 1, 4, 0, 0, 0, 0, 0] [('D', 4), ('D', 5)]
# I [3, 2, 1, 4, 7, 6, 5, 0, 0] []
# D [3, 2, 1, 4, 7, 6, 5, 0, 0] [('D', 7)]
# I [3, 2, 1, 4, 7, 6, 5, 9, 8] []
*/
