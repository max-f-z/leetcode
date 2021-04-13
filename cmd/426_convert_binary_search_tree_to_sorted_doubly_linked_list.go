package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var head *Node
var cur *Node

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}

	head = nil
	cur = nil
	treeToDoublyListHelper(root)

	head.Left = cur
	cur.Right = head

	return head
}

func treeToDoublyListHelper(node *Node) {
	if node != nil {
		treeToDoublyListHelper(node.Left)

		if head == nil {
			head = node
			cur = node
		} else {
			cur.Right = node
			node.Left = cur
			cur = node
		}
		treeToDoublyListHelper(node.Right)
	}
}
