package main

// Intuition

// If the leaf of a tree has 0 coins (an excess of -1 from what it needs), then we should push a coin from its parent onto the leaf. If it has say, 4 coins (an excess of 3), then we should push 3 coins off the leaf. In total, the number of moves from that leaf to or from its parent is excess = Math.abs(num_coins - 1). Afterwards, we never have to consider this leaf again in the rest of our calculation.

// Algorithm

// We can use the above fact to build our answer. Let dfs(node) be the excess number of coins in the subtree at or below this node: namely, the number of coins in the subtree, minus the number of nodes in the subtree. Then, the number of moves we make from this node to and from its children is abs(dfs(node.left)) + abs(dfs(node.right)). After, we have an excess of node.val + dfs(node.left) + dfs(node.right) - 1 coins at this node.

var ans979 int

//lint:ignore U1000 unused
func distributeCoins(root *TreeNode) int {
	ans979 = 0
	distributeCoinsHelper(root)
	return ans979
}

func distributeCoinsHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := distributeCoinsHelper(node.Left)
	r := distributeCoinsHelper(node.Right)

	if l < 0 {
		ans979 -= l
	} else {
		ans979 += l
	}

	if r < 0 {
		ans979 -= r
	} else {
		ans979 += r
	}

	return l + r + node.Val - 1
}
