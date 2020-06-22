package main

import "fmt"

func main() {
	obj := Constructor1()
	obj.Add(1)
	obj.Add(3)
	obj.Add(5)

	fmt.Println(obj.Find(4))
	fmt.Println(obj.Find(7))
}
