package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	return pathSumHelper(root, targetSum, []int{})
}

func pathSumHelper(root *TreeNode, targetSum int, prev []int) [][]int {
	ans := [][]int{}
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			res := make([]int, len(prev)+1)
			copy(res, append(prev, root.Val))
			ans = append(ans, res)
			return ans
		}
		return nil
	}

	if root.Left != nil {
		l := pathSumHelper(root.Left, targetSum-root.Val, append(prev, root.Val))
		if l != nil {
			ans = append(ans, l...)
		}
	}

	if root.Right != nil {
		r := pathSumHelper(root.Right, targetSum-root.Val, append(prev, root.Val))
		if r != nil {
			ans = append(ans, r...)
		}
	}

	return ans
}
