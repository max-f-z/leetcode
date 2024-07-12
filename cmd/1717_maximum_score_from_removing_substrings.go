package main

func maximumGain(s string, x int, y int) int {
	var opt, sub string
	var optScore, subScore int

	if x >= y {
		opt = "ab"
		sub = "ba"
		optScore = x
		subScore = y
	} else {
		opt = "ba"
		sub = "ab"
		optScore = y
		subScore = x
	}

	ans := 0

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == opt[0] && s[i] == opt[1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	ans += optScore * (len(s) - len(stack)) / 2

	stackSub := []byte{}
	for i := 0; i < len(stack); i++ {
		if len(stackSub) > 0 && stackSub[len(stackSub)-1] == sub[0] && stack[i] == sub[1] {
			stackSub = stackSub[:len(stackSub)-1]
		} else {
			stackSub = append(stackSub, stack[i])
		}
	}

	ans += subScore * (len(stack) - len(stackSub)) / 2

	return ans
}
