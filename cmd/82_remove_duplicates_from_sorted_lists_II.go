package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next

	for next != nil && next.Val == head.Val {
		next = next.Next
	}

	if next == head.Next {
		head.Next = deleteDuplicates(next)
		return head
	}

	return deleteDuplicates(next)
}
