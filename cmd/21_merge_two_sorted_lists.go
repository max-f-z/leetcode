package main

//lint:ignore U1000 unused
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

		if l1.Val <= l2.Val {
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

//lint:ignore U1000 unused
func practice21(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
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
		}

		if l2 == nil {
			head.Next = l1
			head = head.Next
			l2 = l1.Next
		}

		if l1.Val <= l2.Val {
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
