package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	head = revList(head)

	cur := head
	tmp := 0
	for cur != nil {
		cur.Val = cur.Val*2 + tmp
		tmp = 0

		if cur.Val >= 10 {
			tmp = cur.Val / 10
		}
		cur.Val = cur.Val % 10

		if cur.Next == nil && tmp != 0 {
			cur.Next = &ListNode{
				Val: tmp,
			}
			break
		}
		cur = cur.Next
	}

	return revList(head)
}

func revList(node *ListNode) *ListNode {
	cur := node
	var prev *ListNode

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	return prev
}
