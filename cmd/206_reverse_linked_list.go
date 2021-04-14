package main

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}
