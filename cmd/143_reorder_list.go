package main

func reorderList(head *ListNode) {
	p1, p2 := head, head

	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	current := p1.Next
	dummy := &ListNode{}

	for current != nil {
		if dummy.Next == nil {
			dummy.Next = current
			current = current.Next
			dummy.Next.Next = nil
		} else {
			tmp := dummy.Next
			dummy.Next = current
			current = current.Next
			dummy.Next.Next = tmp
		}
	}
	p1.Next = nil

	first := head
	second := dummy.Next

	for second != nil {
		tmp := first.Next
		first.Next = second
		second = second.Next
		first.Next.Next = tmp
		first = tmp
	}
}
