package main

//lint:ignore U1000 unused
func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := make([]int, k)
	idx := 0
	cur := head
	curD := head
	for cur != nil {
		if idx == k {
			for i := k - 1; i >= 0; i-- {
				curD.Val = tmp[i]
				curD = curD.Next
			}
			idx = 0
		}

		tmp[idx] = cur.Val
		cur = cur.Next
		idx++
	}

	if idx == k {
		for i := k - 1; i >= 0; i-- {
			curD.Val = tmp[i]
			curD = curD.Next
		}
		idx = 0
	}

	return head
}
