package main

import "fmt"

//lint:ignore U1000 unused
func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return longestConsecutiveHelper(root, 1)
}

func longestConsecutiveHelper(node *TreeNode, max int) int {
	if node == nil {
		return max
	}

	res := max
	if node.Left != nil {
		if node.Left.Val-node.Val == 1 {
			if tmp := longestConsecutiveHelper(node.Left, max+1); tmp > res {
				if tmp == 6 {
					fmt.Println(tmp)
				}
				res = tmp
			}
		} else {
			if tmp := longestConsecutiveHelper(node.Left, 1); tmp > res {
				if tmp == 6 {
					fmt.Println(tmp)
				}
				res = tmp
			}
		}
	}

	if node.Right != nil {
		if node.Right.Val-node.Val == 1 {
			if tmp := longestConsecutiveHelper(node.Right, max+1); tmp > res {
				if tmp == 6 {
					fmt.Println(tmp)
				}
				res = tmp
			}
		} else {
			if tmp := longestConsecutiveHelper(node.Right, 1); tmp > res {
				if tmp == 6 {
					fmt.Println(tmp)
				}
				res = tmp
			}
		}
	}

	return res
}
