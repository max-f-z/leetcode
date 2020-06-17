package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	head := dummy
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil {
			head.Next = l2
			head = head.Next
			l2 = l2.Next
			continue
		}

		if l2 == nil {
			head.Next = l1
			head = head.Next
			l1 = l1.Next
			continue
		}

		if l1.Val <= l2.Val && l1 != nil {
			head.Next = l1
			head = head.Next
			l1 = l1.Next
		} else {
			head.Next = l2
			head = head.Next
			l2 = l2.Next
		}
	}

	return dummy.Next
}
