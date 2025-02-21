package main

type FindElements struct {
	ref map[int]bool
}

func Constructor1261(root *TreeNode) FindElements {
	fe := FindElements{
		ref: map[int]bool{},
	}

	FindElementsDFS(root, 0, fe.ref)

	return fe
}

func FindElementsDFS(node *TreeNode, val int, ref map[int]bool) {
	node.Val = val
	ref[val] = true

	if node.Left != nil {
		FindElementsDFS(node.Left, 2*val+1, ref)
	}

	if node.Right != nil {
		FindElementsDFS(node.Right, 2*val+2, ref)
	}
}

func (fe *FindElements) Find(target int) bool {
	return fe.ref[target]
}
