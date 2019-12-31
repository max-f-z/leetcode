package main

import "fmt"

func main() {
	nums := []int{3, 3, 1, 3}

	fmt.Println(histogram(nums))

	r := [][]byte{
		{'1', '1', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '1', '1'},
	}

	fmt.Println(maximalRectangle(r))
}
