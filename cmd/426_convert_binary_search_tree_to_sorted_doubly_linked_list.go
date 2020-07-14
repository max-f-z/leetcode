package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	sol := Solution{}
	return sol.Run(root)
}

type Solution struct {
	prev       *Node
	head, tail *Node
}

func (s *Solution) Run(root *Node) *Node {
	if root == nil {
		return root
	}
	s.helper(root)
	s.head.Left = s.tail
	s.tail.Right = s.head
	return s.head
}

func (s *Solution) helper(root *Node) {
	if root == nil {
		return
	}
	s.helper(root.Left)
	// find head and tail
	if root.Left == nil && s.head == nil {
		s.head = root
	}
	if root.Right == nil {
		s.tail = root
	}
	// compose linked list
	root.Left = s.prev
	if s.prev != nil {
		s.prev.Right = root
	}
	s.prev = root
	s.helper(root.Right)
}
