package main

func getMaximumGold(grid [][]int) int {
	maxGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			gold := dfsGetMaximumGold(grid, i, j)
			if gold > maxGold {
				maxGold = gold
			}
		}
	}

	return maxGold
}

func dfsGetMaximumGold(grid [][]int, row int, col int) int {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[1]) || grid[row][col] == 0 {
		return 0
	}

	gold := 0

	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	original := grid[row][col]
	grid[row][col] = 0

	for _, direction := range directions {
		subGold := dfsGetMaximumGold(grid, row+direction[0], col+direction[1])
		if subGold > gold {
			gold = subGold
		}
	}

	grid[row][col] = original
	return gold + original
}
