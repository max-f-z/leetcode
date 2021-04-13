package main

import "testing"

func Test_426(t *testing.T) {
	root := &Node{
		Val: 2,
		Right: &Node{
			Val: 3,
		},
	}
	node := treeToDoublyList(root)
	t.Log(node)
}
