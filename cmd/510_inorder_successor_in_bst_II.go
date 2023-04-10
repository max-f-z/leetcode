package main

// Node510 -
type Node510 struct {
	Val    int
	Left   *Node510
	Right  *Node510
	Parent *Node510
}

//lint:ignore U1000 unused
func inorderSuccessor510(node *Node510) *Node510 {
	if node.Right != nil {
		node = node.Right

		for node.Left != nil {
			node = node.Left
		}

		return node
	} else {
		p := node
		successor := new(Node510)

		for p != nil {
			successor = p.Parent
			if successor != nil && p == successor.Left {
				return successor
			}

			p = successor
		}

		return p
	}
}
