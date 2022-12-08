package main

import "testing"

func Test328(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
								},
							},
						},
					},
				},
			},
		},
	}

	ans := oddEvenList(head)

	t.Log(ans)
}
