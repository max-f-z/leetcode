package main

func detectCycle(head *ListNode) *ListNode {
	cols := map[*ListNode]bool{}

	node := head
	for {
		if node == nil {
			return nil
		}

		if cols[node] {

			return node
		}
		cols[node] = true
		node = node.Next
	}
}
