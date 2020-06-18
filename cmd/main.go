package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{Val: -9}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next = &ListNode{Val: 7}
	a := []int{3, 2, 1}
	nextPermutation(a)
	fmt.Println(a)
}
