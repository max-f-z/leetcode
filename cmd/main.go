package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{Val: -9}
	l1.Next = &ListNode{Val: 3}
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 7}

	fmt.Println(mergeKLists([]*ListNode{l1, l2}))
}
