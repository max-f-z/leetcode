package main

import "fmt"

func main() {

	obj := Constructor(2)

	fmt.Println(obj.Get(2))
	obj.Put(2, 6)
	fmt.Println(obj.Get(1))
	obj.Put(1, 5)
	obj.Put(1, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}
