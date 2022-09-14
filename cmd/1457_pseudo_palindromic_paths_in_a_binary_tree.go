package main

func pseudoPalindromicPaths(root *TreeNode) int {
	return pseudoPalindromicPathsHelper(root, map[int]int{})
}

func pseudoPalindromicPathsHelper(node *TreeNode, data map[int]int) int {
	if node == nil {
		return 0
	}

	data[node.Val]++

	if node.Left == nil && node.Right == nil {
		if isPseudoPalindromic(data) {
			data[node.Val]--
			return 1
		}
		data[node.Val]--
		return 0
	}

	l := pseudoPalindromicPathsHelper(node.Left, data)
	r := pseudoPalindromicPathsHelper(node.Right, data)

	data[node.Val]--
	return l + r
}

func isPseudoPalindromic(m map[int]int) bool {
	cnt := 0
	for _, v := range m {
		if v%2 != 0 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}

	return true
}
