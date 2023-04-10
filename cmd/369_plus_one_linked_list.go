package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//lint:ignore U1000 unused
func plusOne(head *ListNode) *ListNode {
	val := plusOneHelper(head, 1)
	if val == 0 {
		return head
	}
	newHead := &ListNode{Val: val}
	newHead.Next = head
	return newHead
}

func plusOneHelper(node *ListNode, add int) (res int) {
	if node.Next == nil {
		node.Val += add
	} else {
		tmp := plusOneHelper(node.Next, add)
		node.Val += tmp
	}

	if node.Val > 9 {
		node.Val, res = node.Val%10, node.Val/10
		return res
	}
	return 0
}
