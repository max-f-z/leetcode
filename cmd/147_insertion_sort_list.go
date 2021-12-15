package main

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}

	current := head

	for current != nil {
		if dummy.Next == nil {
			dummy.Next = current

			tmp := current
			current = current.Next
			tmp.Next = nil
		} else {
			previous := dummy

			for previous.Next != nil && previous.Next.Val < current.Val {
				previous = previous.Next
			}

			if previous.Next == nil {
				previous.Next = current
				tmp := current
				current = current.Next
				tmp.Next = nil
			} else {
				tmp := previous.Next
				previous.Next = current
				current = current.Next
				previous.Next.Next = tmp
			}
		}
	}

	return dummy.Next
}
