package main

import "fmt"

func main() {
	fmt.Println(string(byte(122)))
	fmt.Println(findStrobogrammatic(2))

	fmt.Println(groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}))
}
