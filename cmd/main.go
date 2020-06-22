package main

import "fmt"

func main() {
	b := []byte("the sky is blue")
	reverseWords(b)
	fmt.Println(string(b))
}
