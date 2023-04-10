package main

import "fmt"

//lint:ignore U1000 unused
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

//lint:ignore U1000 unused
func ans(pairs [][]int, a int, b int) ([]int, []int) {
	parents := map[int][]int{}
	for _, v := range pairs {
		if parents[v[1]] == nil || len(parents[v[1]]) == 0 {
			parents[v[1]] = []int{v[0]}
		} else {
			parents[v[1]] = append(parents[v[1]], v[0])
		}
	}

	for k := range parents {
		fmt.Println(k, parents[k])
	}

	fmt.Println(findEarliest(parents, 1))

	return nil, nil
}

func findEarliest(parents map[int][]int, val int) (int, int) {

	if parents[val] != nil && len(parents[val]) != 0 {
		max := 0
		ans := 0
		for _, v := range parents[val] {
			n, newVal := findEarliest(parents, v)
			if n > max {
				max = n
				ans = newVal
			}
		}
		return max + 1, ans
	}

	return 1, val
}
