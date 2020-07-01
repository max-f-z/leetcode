package main

import "math"

func verticalOrder(root *TreeNode) [][]int {
	res := &[][]int{}

	if root == nil {
		return *res
	}

	*res = append(*res, []int{})
	resLeft := &[][]int{}
	stack := &[]*TreeNode{root}
	idxStack := &[]int{0}

	for len(*stack) > 0 {
		var node *TreeNode
		var offset int
		node, offset, (*stack), (*idxStack) = (*stack)[0], (*idxStack)[0], (*stack)[1:], (*idxStack)[1:]

		if offset == 0 {
			(*res)[0] = append((*res)[0], node.Val)
		} else if offset > 0 {
			if len(*res)-1 < offset {
				(*res) = append((*res), []int{node.Val})
			} else {
				(*res)[offset] = append((*res)[offset], node.Val)
			}
		} else {
			if len(*resLeft) < int(math.Abs(float64(offset))) {
				(*resLeft) = append((*resLeft), []int{node.Val})
			} else {
				(*resLeft)[-1-offset] = append((*resLeft)[-1-offset], node.Val)
			}
		}
		if node.Left != nil {
			*stack = append(*stack, node.Left)
			*idxStack = append(*idxStack, offset-1)
		}

		if node.Right != nil {
			*stack = append(*stack, node.Right)
			*idxStack = append(*idxStack, offset+1)
		}
	}

	for i := 0; i < len(*resLeft); i++ {
		*res = prepend(*res, (*resLeft)[i])
	}

	return *res
}

func prepend(x [][]int, y []int) [][]int {
	x = append(x, []int{})
	copy(x[1:], x)
	x[0] = y
	return x
}
