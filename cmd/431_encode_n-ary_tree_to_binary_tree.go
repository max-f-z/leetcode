package main

type Node431 struct {
	Val      int
	Children []*Node431
}

type Codec431 struct {
}

func Constructor431() *Codec431 {
	codec := &Codec431{}
	return codec
}

//lint:ignore ST1006 this
//lint:ignore U1000 unused
func (this *Codec431) encode(root *Node431) *TreeNode {
	return encodeHelper(root)
}

//lint:ignore ST1006 this
//lint:ignore U1000 unused
func (this *Codec431) decode(root *TreeNode) *Node431 {
	return decodeHelper(root)
}

//lint:ignore U1000 unused
func decodeHelper(node *TreeNode) *Node431 {
	if node == nil {
		return nil
	}

	tmp := Node431{Val: node.Val}
	tmp.Children = []*Node431{}

	l := decodeHelper(node.Left)
	if l != nil {
		tmp.Children = append(tmp.Children, l)
		cur := node.Left
		for cur.Right != nil {
			r := decodeHelper(cur.Right)
			tmp.Children = append(tmp.Children, r)
			cur = cur.Right
		}
	}

	return &tmp
}

//lint:ignore U1000 unused
func encodeHelper(node *Node431) *TreeNode {
	if node == nil {
		return nil
	}

	tmp := TreeNode{Val: node.Val}

	if node.Children != nil && len(node.Children) > 0 {
		l := encodeHelper(node.Children[0])
		tmp.Left = l
		cur := l
		for i := 1; i < len(node.Children); i++ {
			rl := encodeHelper(node.Children[i])
			cur.Right = rl
			cur = cur.Right
		}
	}

	return &tmp
}
