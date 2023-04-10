package main

// TreeNode -
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//lint:ignore U1000 unused
type position struct {
	start, end int
}

//lint:ignore U1000 unused
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	cache := make(map[position][]*TreeNode)
	return gen(1, n, cache)
}

func gen(start, end int, cache map[position][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil} // empty tree
	}

	if start == end { // one node
		return []*TreeNode{
			{
				Val: start,
			},
		}
	}

	pos := position{start, end}
	if t, ok := cache[pos]; ok {
		return t
	}

	out := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		ls := gen(start, i-1, cache)
		rs := gen(i+1, end, cache)

		for _, l := range ls {
			for _, r := range rs {
				out = append(out,
					&TreeNode{
						Val:   i,
						Left:  l,
						Right: r,
					})
			}
		}
	}

	cache[pos] = out
	return out
}
