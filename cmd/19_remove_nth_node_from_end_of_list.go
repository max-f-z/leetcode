package main

type ListNode19 struct {
	Val  int
	Next *ListNode19
}

//lint:ignore U1000 unused
func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {
	dummy := &ListNode19{Next: head}
	p1, p2 := dummy, dummy

	for i := 0; i <= n; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next
}
