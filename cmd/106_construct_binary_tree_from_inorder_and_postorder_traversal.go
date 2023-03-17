package main

// Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// Output: [3,9,20,null,null,15,7]

type treeBuilder struct {
	inorder   []int
	postorder []int
	idx       map[int]int
	postIdx   int
}

func buildTree106(inorder []int, postorder []int) *TreeNode {
	idx := map[int]int{}
	for i := 0; i < len(inorder); i++ {
		idx[inorder[i]] = i
	}

	tb := &treeBuilder{
		inorder:   inorder,
		postorder: postorder,
		idx:       idx,
		postIdx:   len(inorder) - 1,
	}

	return tb.buildTreeHelper(0, len(inorder)-1)
}

func (tb *treeBuilder) buildTreeHelper(l, r int) *TreeNode {
	if l > r || tb.postIdx < 0 {
		return nil
	}

	val := tb.postorder[tb.postIdx]
	tb.postIdx--

	valIdx := tb.idx[val]
	right := tb.buildTreeHelper(valIdx+1, r)
	left := tb.buildTreeHelper(l, valIdx-1)

	node := &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}

	return node
}
