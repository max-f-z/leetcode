package main

import "testing"

func Test142(t *testing.T) {
	root := &ListNode{
		Val: 3,
	}
	root1 := &ListNode{
		Val: 2,
	}

	root2 := &ListNode{
		Val: 0,
	}

	root3 := &ListNode{
		Val: -4,
	}

	root.Next = root1
	root1.Next = root2
	root2.Next = root3
	root3.Next = root1

	n := detectCycle(root)
	t.Log(n)
}
