package main

func swapPairs(head *ListNode) *ListNode {
	flag := true
	cur := head
	for cur != nil {
		if flag && cur.Next != nil {
			cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
		}
		cur = cur.Next
		flag = !flag
	}

	return head
}
