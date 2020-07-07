package main

import (
	"fmt"
	"math"
)

func main() {
	root := &ListNode{Val: 9}
	root.Next = &ListNode{Val: 9}
	root.Next.Next = &ListNode{Val: 9}

	fmt.Println(plusOne(root))
}

func powerTrue(k int) bool {
	for i := 0; i < 100; i++ {
		if k == int(math.Pow(2, float64(i))) {
			return true
		}
	}

	return false
}
