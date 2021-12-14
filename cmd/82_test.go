package main

import "testing"

func Test82(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}

	head.Next = &ListNode{
		Val: 2,
	}

	head.Next.Next = &ListNode{
		Val: 3,
	}

	head.Next.Next.Next = &ListNode{
		Val: 3,
	}

	head.Next.Next.Next.Next = &ListNode{
		Val: 4,
	}

	head.Next.Next.Next.Next.Next = &ListNode{
		Val: 4,
	}

	head.Next.Next.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}

	ans := deleteDuplicates(head)

	t.Log(ans)
}
