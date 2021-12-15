package main

import "testing"

func Test147(t *testing.T) {
	head := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}

	ans := insertionSortList(head)

	t.Log(ans)
}
