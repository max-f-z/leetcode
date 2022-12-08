package main

func oddEvenList(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	p1, p2 := head, head.Next

	evenHead := p2

	for p2 != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = evenHead

	return head
}
