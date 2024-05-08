package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test2816(t *testing.T) {
	head := &ListNode{
		Val: 5,
		// Next: &ListNode{
		// 	Val: 9,
		// 	Next: &ListNode{
		// 		Val: 9,
		// 	},
		// },
	}
	head = doubleIt(head)

	assert.Check(t, head != nil)
}
