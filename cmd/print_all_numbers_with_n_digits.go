package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printNum(prefix string, n int) {
	if n == 1 {
		for i := 0; i < 10; i++ {
			str := fmt.Sprintf("%s%d ", prefix, i)
			str = strings.TrimLeft(str, "0")
			if str == " " {
				continue
			}
			fmt.Print(str)
		}
		fmt.Print("\n")
		return
	}

	for i := 0; i < 10; i++ {
		printNum(prefix+strconv.Itoa(i), n-1)
	}
}
