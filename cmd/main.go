package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(alienOrder([]string{"bsusz", "rhn", "gfbrwec", "kuw", "qvpxbexnhx", "gnp", "laxutz", "qzxccww"}))

	// var cur, root *TreeNode
	// cur = nil
	// stack := []*TreeNode{}

	// for i := 0; i < len(arr); i++ {
	// 	tmp := TreeNode{}
	// 	tmp.Val = arr[i]
	// 	if i == 0 {
	// 		root = &tmp
	// 	}

	// 	if cur == nil {
	// 		cur = &tmp
	// 	} else if cur.Left == nil {
	// 		cur.Left = &tmp
	// 		stack = append(stack, &tmp)
	// 	} else if cur.Right == nil {
	// 		cur.Right = &tmp
	// 		stack = append(stack, &tmp)
	// 	} else {
	// 		cur, stack = stack[0], stack[1:]
	// 		cur.Left = &tmp
	// 		stack = append(stack, &tmp)
	// 	}
	// }
	// obj := Constructor307([]int{9, -8})
	// obj.Update(0, 3)
	// fmt.Println(obj.SumRange(1, 1))
	// fmt.Println(obj.SumRange(0, 1))
	// obj.Update(1, -3)
	// fmt.Println(obj.SumRange(0, 1))

	root := &TreeNode{Val: 2}
	// root.Left = &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	// root.Left.Left = &TreeNode{Val: 1}
	// root.Left.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 1}

	fmt.Println(largestBSTSubtree(root))
}

func powerTrue(k int) bool {
	for i := 0; i < 100; i++ {
		if k == int(math.Pow(2, float64(i))) {
			return true
		}
	}

	return false
}
