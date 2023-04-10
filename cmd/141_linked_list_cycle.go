package main

//lint:ignore U1000 unused
func hasCycle(head *ListNode) bool {
	p1, p2 := head, head

	for p1 != nil && p2 != nil {
		if p2.Next == nil || p2.Next.Next == nil {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next.Next

		if p1 == p2 {
			return true
		}
	}

	return false
}
