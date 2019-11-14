package main

import "fmt"

func main() {
	test := [][]int{[]int{1, 3, 1}, []int{1, 5, 1}, []int{4, 2, 1}}

	fmt.Println(minPathSum(test))
}
