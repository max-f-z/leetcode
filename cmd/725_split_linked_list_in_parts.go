package main

//lint:ignore U1000 unused
func splitListToParts(head *ListNode, k int) []*ListNode {
	cur := head
	cnt := 0

	for cur != nil {
		cnt++
		cur = cur.Next
	}

	avg := cnt / k
	extra := cnt % k

	ans := make([]*ListNode, k)

	cur = head
	idx := 0
	var prev *ListNode
	ans[0] = head
	ansIdx := 1
	for cur != nil {
		if idx < avg {
			prev = cur
			cur = cur.Next
			idx++
			continue
		}

		if idx == avg && extra > 0 {
			prev = cur
			cur = cur.Next
			extra--
			idx++
			continue
		}

		prev.Next = nil
		ans[ansIdx] = cur
		ansIdx++
		idx = 0
	}

	return ans
}
