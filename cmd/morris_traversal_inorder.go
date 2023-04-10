package main

//lint:ignore U1000 unused
func morrisInorder(root *TreeNode) []int {
	ans := []int{}

	current := root

	for current != nil {
		guide := current.Left

		if guide == nil {
			ans = append(ans, current.Val)
			current = current.Right
		} else {
			for guide.Right != nil && guide.Right != current {
				guide = guide.Right
			}

			if guide.Right == nil {
				guide.Right = current
				current = current.Left
			} else {
				guide.Right = nil
				ans = append(ans, current.Val)
				current = current.Right
			}
		}
	}

	return ans
}
