package main

//lint:ignore U1000 unused
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	cur := prev.Next

	for i := left; i < right; i++ {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = prev.Next
		prev.Next = tmp
	}

	return dummy.Next
}
