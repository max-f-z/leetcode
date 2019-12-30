package main

import "fmt"

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
}
