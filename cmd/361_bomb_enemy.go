package main

func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	h, w := len(grid), len(grid[0])
	left, right, up, down := make([][]int, h), make([][]int, h), make([][]int, h), make([][]int, h)
	for i := 0; i < h; i++ {
		left[i], right[i], up[i], down[i] = make([]int, w), make([]int, w), make([]int, w), make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 'W' {
				left[i][j] = 0
				up[i][j] = 0
				continue
			}

			if grid[i][j] == 'E' {
				if j < w-1 {
					left[i][j+1] = left[i][j] + 1
				}

				if i < h-1 {
					up[i+1][j] = up[i][j] + 1
				}
			} else {
				if j < w-1 {
					left[i][j+1] = left[i][j]
				}

				if i < h-1 {
					up[i+1][j] = up[i][j]
				}
			}

		}
	}

	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			if grid[i][j] == 'W' {
				right[i][j] = 0
				down[i][j] = 0
				continue
			}

			if grid[i][j] == 'E' {
				if j > 0 {
					right[i][j-1] = right[i][j] + 1
				}

				if i > 0 {
					down[i-1][j] = down[i][j] + 1
				}
			} else {
				if j > 0 {
					right[i][j-1] = right[i][j]
				}

				if i > 0 {
					down[i-1][j] = down[i][j]
				}
			}
		}
	}

	max := 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '0' {
				if max < (right[i][j] + left[i][j] + up[i][j] + down[i][j]) {
					max = right[i][j] + left[i][j] + up[i][j] + down[i][j]
				}
			}
		}
	}

	return max
}
