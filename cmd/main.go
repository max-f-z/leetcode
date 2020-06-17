package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{Val: -9}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next = &ListNode{Val: 7}

	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
